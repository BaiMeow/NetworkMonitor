package bgp

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/osrg/gobgp/v3/pkg/server"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"math"
	"net/netip"
	"slices"
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

		bgpServer := server.NewBgpServer()
		go bgpServer.Serve()
		if err := bgpServer.StartBgp(context.Background(), &apipb.StartBgpRequest{
			Global: &apipb.Global{
				Asn:        asn,
				RouterId:   routerIdRaw,
				ListenPort: int32(port),
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
						Enabled:     true,
						MultihopTtl: 32,
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
						Enabled:     true,
						MultihopTtl: 32,
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
	s *server.BgpServer
}

func (f *BGP) GetData() ([]byte, error) {
	// Wait ESTABLISHED
	for i := 0; i < 10; i++ {
		var established bool
		f.s.ListPeer(context.Background(), &apipb.ListPeerRequest{}, func(peer *apipb.Peer) {
			if peer.State.SessionState == apipb.PeerState_ESTABLISHED {
				established = true
				if i != 0 {
					fmt.Println("BGP Session State:", peer.State.SessionState)
				}
			} else {
				fmt.Println("BGP Session State:", peer.State.SessionState)
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

	var paths [][]uint32
	if err := f.s.ListPath(context.Background(), &apipb.ListPathRequest{
		Family: &apipb.Family{
			Afi:  apipb.Family_AFI_IP,
			Safi: apipb.Family_SAFI_UNICAST,
		},
	}, func(destination *apipb.Destination) {
		for _, p := range destination.GetPaths() {
			idx := slices.IndexFunc(p.Pattrs, func(a *anypb.Any) bool {
				return a.GetTypeUrl() == "type.googleapis.com/apipb.AsPathAttribute"
			})
			if idx == -1 {
				continue
			}
			asPathAttrPb := p.Pattrs[idx]
			var asPathAttr apipb.AsPathAttribute
			if err := asPathAttrPb.UnmarshalTo(&asPathAttr); err != nil {
				log.Println("unmarshal ASPathAttr failed:", err)
				continue
			}
			for _, se := range asPathAttr.Segments {
				paths = append(paths, se.Numbers)
			}
		}
	}); err != nil {
		return nil, err
	}
	fmt.Println(len(paths))
	return nil, nil
}
