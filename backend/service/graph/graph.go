package graph

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	parse2 "github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"log"
	"net/netip"
	"slices"
)

func GetOSPF(asn uint32) *parse2.OSPF { return graph.OSPF[asn] }

func GetBGP() *parse2.BGP {
	bgp := *graph.BGP
	recordASNs, err := uptime.AllASNRecord()
	if err != nil {
		log.Println("read uptime fail:", err)
		return &bgp
	}

	currentASs := make([]*parse2.AS, len(bgp.AS))
	copy(currentASs, bgp.AS)
	var addon []*parse2.AS

	slices.SortFunc(currentASs, func(i, j *parse2.AS) int {
		if i.ASN < j.ASN {
			return -1
		} else if i.ASN > j.ASN {
			return 1
		} else {
			return 0
		}
	})

	// append ghost AS
	for _, as := range recordASNs {
		if _, found := slices.BinarySearchFunc(currentASs, as, func(as *parse2.AS, as2 uint32) int {
			if as.ASN < as2 {
				return -1
			} else if as.ASN > as2 {
				return 1
			} else {
				return 0
			}
		}); !found {
			addon = append(addon, &parse2.AS{ASN: as, Network: []netip.Prefix{}})
		}
	}

	bgp.AS = append(currentASs, addon...)
	return &bgp
}

func ListAvailable() []map[string]any {
	var ls []map[string]any
	if len(graph.BGP.AS) != 0 {
		ls = append(ls, map[string]any{
			"type": "bgp",
		})
	}

	for asn := range graph.OSPF {
		ls = append(ls, map[string]any{
			"type": "ospf",
			"asn":  asn,
		})
	}

	return ls
}
