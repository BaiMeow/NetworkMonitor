package bgp

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	apipb "github.com/osrg/gobgp/v3/api"
	gobgplog "github.com/osrg/gobgp/v3/pkg/log"
	"github.com/osrg/gobgp/v3/pkg/server"
	"github.com/pkg/errors"
	"log"
	"math"
	"net/netip"
	"time"
)

func init() {
	fetch.Register("bgp", func(m map[string]any) (fetch.Fetcher, error) {
		asnRaw, ok := m["asn"].(int)
		if !ok {
			return nil, fmt.Errorf("asn is not int")
		}
		if asnRaw > math.MaxUint32 || asnRaw < 0 {
			return nil, fmt.Errorf("invalid asn number: %d", asnRaw)
		}
		asn := uint32(asnRaw)

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

		logger := gobgplog.NewDefaultLogger()
		logger.SetLevel(gobgplog.ErrorLevel)
		bgpServer := server.NewBgpServer(server.LoggerOption(logger))
		go bgpServer.Serve()
		if err := bgpServer.StartBgp(context.Background(), &apipb.StartBgpRequest{
			Global: &apipb.Global{
				Asn:              asn,
				RouterId:         routerIdRaw,
				ListenPort:       int32(port),
				UseMultiplePaths: true,
			},
		}); err != nil {
			return nil, err
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
				return nil, err
			}
			if err := bgpServer.AddDynamicNeighbor(context.Background(), &apipb.AddDynamicNeighborRequest{
				DynamicNeighbor: &apipb.DynamicNeighbor{
					Prefix:    "0.0.0.0/0",
					PeerGroup: "route-collector",
				},
			}); err != nil {
				return nil, err
			}
		} else {
			peerASN, ok := m["peer-asn"].(int)
			if !ok {
				return nil, fmt.Errorf("peer-asn is not int")
			}
			if peerASN > math.MaxUint32 || peerASN < 0 {
				return nil, fmt.Errorf("invalid peer-asn number: %d", asnRaw)
			}

			neighborAddr, ok := m["neighbor-addr"].(string)
			if !ok {
				return nil, fmt.Errorf("neighbor-addr is not string")
			}
			if _, err := netip.ParseAddr(neighborAddr); err != nil {
				return nil, fmt.Errorf("invalid neighbor-addr: %s", neighborAddr)
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
				return nil, err
			}
		}

		return &BGP{
			s: bgpServer,
		}, nil
	})
}

type BGP struct {
	fetch.Base
	s *server.BgpServer
}

func (f *BGP) GetData() (any, error) {
	// Wait ESTABLISHED
	for i := 0; i < 10; i++ {
		var established bool
		f.s.ListPeer(context.Background(), &apipb.ListPeerRequest{}, func(peer *apipb.Peer) {
			if peer.State.SessionState == apipb.PeerState_ESTABLISHED {
				established = true
				if i != 0 {
					log.Println("BGP Session State:", peer.State.SessionState)
				}
			} else {
				log.Println("BGP Session State:", peer.State.SessionState)
			}
		})
		if established {
			break
		}
		if i != 9 {
			time.Sleep(time.Second * 3)
			continue
		}
		return nil, errors.New("BGP session failed")
	}

	var destinations []*apipb.Destination
	if err := f.s.ListPath(context.Background(), &apipb.ListPathRequest{
		Family: &apipb.Family{
			Afi:  apipb.Family_AFI_IP,
			Safi: apipb.Family_SAFI_UNICAST,
		},
		EnableFiltered: true,
	}, func(destination *apipb.Destination) {
		destinations = append(destinations, destination)
	}); err != nil {
		return nil, err
	}
	return destinations, nil
}

func (f *BGP) Stop() error {
	f.s.StopBgp(context.Background(), nil)
	f.s.Stop()
	return nil
}
