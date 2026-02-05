package bgp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v4/api"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func init() {
	fetch.Register("gobgp", func(m map[string]any) (fetch.Fetcher, error) {
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
		}

		client, err := grpc.NewClient(target, opts...)
		if err != nil {
			return nil, fmt.Errorf("new grpc client fail: %v", err)
		}
		return &GoBGP{
			api: apipb.NewGoBgpServiceClient(client),
		}, nil
	})
}

type GoBGP struct {
	fetch.Base
	api apipb.GoBgpServiceClient
}

func (f *GoBGP) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/bgp/GoBGP.GetData",
	)
	defer span.End()

	var estCount int
	{
		resp, err := recvResponse(f.api.ListPeer(ctx, &apipb.ListPeerRequest{}))
		if err != nil {
			return nil, fmt.Errorf("list bgp peer: %v", err)
		}
		for _, v := range resp {
			if v.Peer.State.SessionState == apipb.PeerState_SESSION_STATE_ESTABLISHED {
				estCount++
			}
		}
		span.SetAttributes(attribute.Int("established_count", estCount))
		if estCount == 0 {
			return nil, errors.New("no established peer")
		}
	}

	var paths []*apipb.Path
	{
		resp, err := recvResponse(f.api.ListPath(ctx, &apipb.ListPathRequest{
			TableType: apipb.TableType_TABLE_TYPE_GLOBAL,
			Family: &apipb.Family{
				Afi:  apipb.Family_AFI_IP,
				Safi: apipb.Family_SAFI_UNICAST,
			},
			EnableFiltered: true,
		}))
		if err != nil {
			return nil, fmt.Errorf("list bgp path: %v", err)
		}
		for _, v := range resp {
			paths = append(paths, v.GetDestination().GetPaths()...)
		}
	}
	span.SetAttributes(attribute.Int("path_count", len(paths)))
	return paths, nil
}

func countPath(dess []*apipb.Destination) int {
	var sum int
	for _, des := range dess {
		sum += len(des.GetPaths())
	}
	return sum
}

type Revivable[T any] interface {
	Recv() (T, error)
}

func recvResponse[T any](c Revivable[T], err error) ([]T, error) {
	if err != nil {
		return nil, err
	}
	var arr []T
	for {
		v, err := c.Recv()
		if err != nil {
			if err == io.EOF {
				return arr, nil
			} else {
				return nil, err
			}
		}
		arr = append(arr, v)
	}
}
