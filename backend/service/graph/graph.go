package graph

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"log"
	"net/netip"
	"slices"
)

func GetOSPF(asn uint32) *parse.OSPF {
	return graph.GetOSPF(asn)
}

func GetBGP() *parse.BGP {
	bgp := *graph.GetBGP()
	recordASNs, err := uptime.AllASNRecord()
	if err != nil {
		log.Println("read uptime fail:", err)
		return &bgp
	}

	currentASs := make([]*parse.AS, len(bgp.AS))
	copy(currentASs, bgp.AS)
	var addon []*parse.AS

	slices.SortFunc(currentASs, func(i, j *parse.AS) int {
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
		if _, found := slices.BinarySearchFunc(currentASs, as, func(as *parse.AS, as2 uint32) int {
			if as.ASN < as2 {
				return -1
			} else if as.ASN > as2 {
				return 1
			} else {
				return 0
			}
		}); !found {
			addon = append(addon, &parse.AS{ASN: as, Network: []netip.Prefix{}})
		}
	}

	bgp.AS = append(currentASs, addon...)
	return &bgp
}

func ListAvailable() []map[string]any {
	var ls []map[string]any
	if len(graph.GetBGP().AS) != 0 {
		ls = append(ls, map[string]any{
			"type": "bgp",
		})
	}

	for asn := range graph.GetAllOSPF() {
		ls = append(ls, map[string]any{
			"type": "ospf",
			"asn":  asn,
		})
	}

	return ls
}
