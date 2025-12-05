package bgp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"math"
	"sync"
	"time"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v3/api"
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
		api := apipb.NewGobgpApiClient(client)

		ctx, cancel := context.WithCancel(context.Background())
		fetcher := &GoBGPWatch{
			cancel: cancel,
			ctx:    ctx,
			api:    api,
		}
		go fetcher.Run()
		return fetcher, nil
	})
}

type GoBGPWatch struct {
	fetch.Base
	api    apipb.GobgpApiClient
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
	span.SetAttributes(attribute.Int("destinations", len(paths)))
	return paths, nil
}

func (f *GoBGPWatch) Run() {
	for {
		ctx, cancel := context.WithTimeout(f.ctx, time.Hour)
		res, err := f.api.WatchEvent(ctx, &apipb.WatchEventRequest{
			Table: &apipb.WatchEventRequest_Table{
				Filters: []*apipb.WatchEventRequest_Table_Filter{
					{
						Type: apipb.WatchEventRequest_Table_Filter_POST_POLICY,
						Init: true,
					},
				},
			},
		})
		if err != nil {
			cancel()
			if errors.Is(err, context.Canceled) {
				return
			}

			log.Println("watch err:", err)

			// clear data
			f.lock.Lock()
			f.paths = nil
			f.lock.Unlock()

			// wait and retry
			time.Sleep(time.Second)
			continue
		}
		for {
			event, err := res.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Println("watch event EOF")
				} else {
					log.Printf("watch event error: %v", err)
				}
				break
			}
			paths := event.GetTable().GetPaths()
			f.lock.Lock()
			if f.paths == nil {
				f.paths = make(map[string]*apipb.Path)
			}
			for _, p := range paths {
				var prefix apipb.IPAddressPrefix
				err := p.GetNlri().UnmarshalTo(&prefix)
				if err != nil {
					fmt.Println("unmarshal prefix error:", err)
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
		}
		cancel()
	}
}
