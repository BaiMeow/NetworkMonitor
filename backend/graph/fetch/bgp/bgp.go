package bgp

import (
	"context"
	"fmt"
	"log"
	"net/netip"
	"time"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"github.com/BaiMeow/NetworkMonitor/utils"
	apipb "github.com/osrg/gobgp/v4/api"
	"github.com/osrg/gobgp/v4/pkg/apiutil"
	"github.com/osrg/gobgp/v4/pkg/packet/bgp"
	"github.com/osrg/gobgp/v4/pkg/server"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/attribute"
)

func init() {
	fetch.Register("bgp", func(m map[string]any) (fetch.Fetcher, error) {
		asn, err := utils.MustASN(m["asn"])
		if err != nil {
			return nil, fmt.Errorf("asn is not valid asn: %v", err)
		}

		routerIdRaw, ok := m["router-id"].(string)
		if !ok {
			return nil, fmt.Errorf("router-id is not string")
		}
		if _, err := netip.ParseAddr(routerIdRaw); err != nil {
			return nil, fmt.Errorf("invalid router-id: %v", err)
		}

		mode, ok := m["mode"].(string)
		if !ok {
			return nil, fmt.Errorf("mode is not string")
		}

		var port int
		switch mode {
		case "listen":
			if m["port"] == nil {
				port = 179
				break
			}
			port, ok = m["port"].(int)
			if !ok {
				return nil, fmt.Errorf("port not int")
			}
		case "dial":
			port = -1
		default:
			return nil, fmt.Errorf("invalid bgp fetcher mode: %s", mode)
		}
		bgpServer := server.NewBgpServer()
		ctx, cancel := context.WithCancel(context.Background())
		bgp := &BGP{
			s:      bgpServer,
			cancel: cancel,
		}
		go bgpServer.Serve()
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				if err := bgpServer.StartBgp(ctx, &apipb.StartBgpRequest{
					Global: &apipb.Global{
						Asn:              asn,
						RouterId:         routerIdRaw,
						ListenPort:       int32(port),
						UseMultiplePaths: true,
					},
				}); err != nil {
					log.Printf("start bgp fail: %v", err)
					time.Sleep(time.Second * 3)
					continue
				}
				break
			}

			if mode == "listen" {
				if err := bgpServer.AddPeerGroup(context.Background(), &apipb.AddPeerGroupRequest{
					PeerGroup: &apipb.PeerGroup{
						Conf: &apipb.PeerGroupConf{
							PeerGroupName: "route-collector",
							LocalAsn:      asn,
						},
						EbgpMultihop: &apipb.EbgpMultihop{
							Enabled: true,
						},
						AfiSafis: []*apipb.AfiSafi{{
							Config: &apipb.AfiSafiConfig{
								Family: &apipb.Family{
									Afi:  apipb.Family_AFI_IP,
									Safi: apipb.Family_SAFI_UNICAST,
								},
							},
							AddPaths: &apipb.AddPaths{
								Config: &apipb.AddPathsConfig{
									Receive: true,
								},
							},
						}},
					},
				}); err != nil {
					log.Printf("add peer group fail: %v", err)
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
					return
				}
				if err := bgpServer.AddDynamicNeighbor(context.Background(), &apipb.AddDynamicNeighborRequest{
					DynamicNeighbor: &apipb.DynamicNeighbor{
						Prefix:    "0.0.0.0/0",
						PeerGroup: "route-collector",
					},
				}); err != nil {
					log.Printf("add dyn neighbor: %v", err)
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
					return
				}
			} else {
				peerASN, err := utils.MustASN(m["peer-asn"])
				if err != nil {
					log.Printf("peer-asn is not valid asn: %v", err)
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
					return
				}
				neighborAddr, ok := m["neighbor-addr"].(string)
				if !ok {
					log.Printf("neighbor-addr is not string")
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
					return
				}
				if _, err := netip.ParseAddr(neighborAddr); err != nil {
					log.Printf("invalid neighbor-addr: %s", neighborAddr)
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
					return
				}
				if err := bgpServer.AddPeer(context.Background(), &apipb.AddPeerRequest{
					Peer: &apipb.Peer{
						Conf: &apipb.PeerConf{
							PeerAsn:         uint32(peerASN),
							NeighborAddress: neighborAddr,
						},
						EbgpMultihop: &apipb.EbgpMultihop{
							Enabled: true,
						},
						AfiSafis: []*apipb.AfiSafi{{
							Config: &apipb.AfiSafiConfig{
								Family: &apipb.Family{
									Afi:  apipb.Family_AFI_IP,
									Safi: apipb.Family_SAFI_UNICAST,
								},
							},
							AddPaths: &apipb.AddPaths{
								Config: &apipb.AddPathsConfig{
									Receive: true,
								},
							},
						}},
					},
				}); err != nil {
					log.Printf("add peer fail: %s", neighborAddr)
					if err := bgp.CleanUp(); err != nil {
						log.Printf("cleaup: %v", err)
					}
				}
			}
		}()

		return bgp, nil
	})
}

type BGP struct {
	fetch.Base
	s      *server.BgpServer
	cancel context.CancelFunc
}

func (f *BGP) GetData(ctx context.Context) (any, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"fetch/bgp/BGP.GetData",
	)
	defer span.End()

	var estCount int
	if err := f.s.ListPeer(ctx, &apipb.ListPeerRequest{}, func(peer *apipb.Peer) {
		if peer.State.SessionState == apipb.PeerState_SESSION_STATE_ESTABLISHED {
			estCount++
		}
	}); err != nil {
		return nil, fmt.Errorf("list bgp peer: %v", err)
	}
	span.SetAttributes(attribute.Int("established_count", estCount))
	if estCount == 0 {
		return nil, errors.New("no established peer")
	}

	var destinations [][2]any
	if err := f.s.ListPath(apiutil.ListPathRequest{
		TableType:      apipb.TableType_TABLE_TYPE_GLOBAL,
		Family:         bgp.NewFamily(bgp.AFI_IP, bgp.SAFI_UNICAST),
		EnableFiltered: true,
	}, func(prefix bgp.NLRI, paths []*apiutil.Path) {
		destinations = append(destinations, [2]any{prefix, paths})
	}); err != nil {
		return nil, err
	}
	// context timeout may interrupt ListPath without any err reported, which leads to incomplete data returning.
	// so if we find it timeout, treat data as broken and return err.
	select {
	case <-ctx.Done():
		return nil, context.Cause(ctx)
	default:
	}
	span.SetAttributes(attribute.Int("destination_count", len(destinations)))
	//span.SetAttributes(attribute.Int("path_count", countPath(destinations)))
	return destinations, nil
}

func (f *BGP) CleanUp() error {
	if f.cancel != nil {
		f.cancel()
	}
	err := f.s.StopBgp(context.Background(), &apipb.StopBgpRequest{})
	if err != nil {
		log.Printf("stop bgp: %v", err)
	}
	f.s.Stop()
	return nil
}

func countPath(dess []*apipb.Destination) int {
	var sum int
	for _, des := range dess {
		sum += len(des.GetPaths())
	}
	return sum
}
