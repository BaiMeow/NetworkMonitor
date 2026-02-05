package bgp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v4/api"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	fetch.Register("gobgp-watch", func(m map[string]any) (fetch.Fetcher, error) {
		target, ok := m["target"].(string)
		if !ok {
			return nil, fmt.Errorf("target is not string")
		}

		var opts []grpc.DialOption
		if tlsConfig, ok := m["mtls"].(map[string]any); ok {
			caCert, ok := tlsConfig["ca"].(string)
			if !ok {
				return nil, fmt.Errorf("cacert is not string")
			}
			clientCert, ok := tlsConfig["cert"].(string)
			if !ok {
				return nil, fmt.Errorf("client cert is not string")
			}
			clientKey, ok := tlsConfig["key"].(string)
			if !ok {
				return nil, fmt.Errorf("client key is not string")
			}
			rootCertPool := x509.NewCertPool()
			rootCertPool.AppendCertsFromPEM([]byte(caCert))
			pair, err := tls.X509KeyPair([]byte(clientCert), []byte(clientKey))
			if err != nil {
				return nil, fmt.Errorf("new tls client fail: %v", err)
			}
			opts = append(opts, grpc.WithTransportCredentials(
				credentials.NewTLS(&tls.Config{
					RootCAs:      rootCertPool,
					Certificates: []tls.Certificate{pair},
				}),
			))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)))

		client, err := grpc.NewClient(target, opts...)
		if err != nil {
			return nil, fmt.Errorf("new grpc client fail: %v", err)
		}
		api := apipb.NewGoBgpServiceClient(client)

		ctx, cancel := context.WithCancel(context.Background())
		fetcher := &GoBGPWatch{
			cancel: cancel,
			ctx:    ctx,
			api:    api,
		}
		go fetcher.run()
		return fetcher, nil
	})
}

type GoBGPWatch struct {
	fetch.Base
	api    apipb.GoBgpServiceClient
	ctx    context.Context
	cancel context.CancelFunc

	lock  sync.Mutex
	paths map[string]*apipb.Path
}

func (f *GoBGPWatch) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/bgp/GoBGPWatch.GetData",
	)
	defer span.End()
	var paths []*apipb.Path
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.paths == nil {
		return nil, errors.New("GoBGP API disconnected")
	}
	for _, v := range f.paths {
		paths = append(paths, v)
	}
	span.SetAttributes(attribute.Int("paths", len(paths)))
	return paths, nil
}

func (f *GoBGPWatch) CleanUp() error {
	f.cancel()
	return nil
}

const fetchALLInterval = time.Hour

func (f *GoBGPWatch) run() {
	session, cancel, err := f.mustConnect()
	if err != nil {
		slog.Error("GoBGP watch mustConnect fail", "err", err)
		return
	}
	var (
		lock          sync.Mutex
		init          = true
		fetchALLTimer = time.NewTimer(fetchALLInterval)
	)

	// watch loop
	go func() {
		for {
			if f.ctx.Err() != nil {
				return
			}
			lock.Lock()
			err := f.watch(session, init)
			if init {
				init = false
			}
			if err != nil && !errors.Is(err, context.Canceled) {
				slog.Error("GoBGP watch fail", "err", err)
				fetchALLTimer.Stop()
				cancel()
				// reconnect
				newSession, newCancel, err := f.mustConnect()
				if err != nil {
					slog.Error("GoBGP watch mustConnect fail", "err", err)
					return
				}
				session = newSession
				cancel = newCancel
				init = true
				fetchALLTimer.Reset(fetchALLInterval)
			}
			lock.Unlock()
		}
	}()

	// reconnect
	for {
		select {
		case <-f.ctx.Done():
			cancel()
			return
		case <-time.After(fetchALLInterval):
			fmt.Println("GoBGP watch reconnect")
			newSession, newCancel, err := f.mustConnect()
			if err != nil {
				slog.Error("GoBGP watch mustConnect fail", "err", err)
				cancel()
				return
			}
			cancel()
			lock.Lock()
			session = newSession
			cancel = newCancel
			init = true
			lock.Unlock()
		}
	}
}

func (f *GoBGPWatch) mustConnect() (apipb.GoBgpService_WatchEventClient, context.CancelFunc, error) {
	for {
		// try mustConnect
		ctx, cancel := context.WithCancel(f.ctx)
		res, err := f.api.WatchEvent(ctx, &apipb.WatchEventRequest{
			Table: &apipb.WatchEventRequest_Table{
				Filters: []*apipb.WatchEventRequest_Table_Filter{
					{
						Type: apipb.WatchEventRequest_Table_Filter_TYPE_POST_POLICY,
						Init: true,
					},
				},
			},
		})
		if err != nil {
			cancel()
			// cancel by caller
			if f.ctx.Err() != nil {
				return nil, nil, f.ctx.Err()
			}
			slog.Error("watch err:", "err", err)
			// wait and retry
			time.Sleep(time.Second)
			continue
		}
		return res, cancel, nil
	}
}

func (f *GoBGPWatch) watch(c apipb.GoBgpService_WatchEventClient, init bool) error {
	event, err := c.Recv()
	if err != nil {
		return err
	}
	paths := event.GetTable().GetPaths()
	f.lock.Lock()
	if init {
		f.paths = make(map[string]*apipb.Path)
	}
	for _, p := range paths {
		prefix := p.GetNlri().GetPrefix()
		if prefix == nil {
			continue
		}
		key := fmt.Sprintf("%s/%d|%s|%d", prefix.GetPrefix(), prefix.GetPrefixLen(), p.SourceId, p.GetIdentifier())
		if p.IsWithdraw {
			if f.paths[key] == nil {
				fmt.Println("delete non existed path: ", key)
			} else {
				delete(f.paths, key)
			}
		} else {
			f.paths[key] = p
		}
	}
	f.lock.Unlock()
	return nil
}
