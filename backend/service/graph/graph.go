package graph

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"log"
	"net/netip"
	"slices"
)

func GetOSPF(asn uint32) *entity.OSPF {
	return graph.GetOSPF(asn)
}

func GetBGP() *entity.BGP {
	bgp := graph.GetBGP()
	if bgp == nil {
		return nil
	}
	recordASNs, err := uptime.AllASNRecord()
	if err != nil {
		log.Println("read uptime fail:", err)
		return bgp
	}

	currentASs := make([]*entity.AS, len(bgp.AS))
	copy(currentASs, bgp.AS)
	var addon []*entity.AS

	slices.SortFunc(currentASs, func(i, j *entity.AS) int {
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
		if _, found := slices.BinarySearchFunc(currentASs, as, func(as *entity.AS, as2 uint32) int {
			if as.ASN < as2 {
				return -1
			} else if as.ASN > as2 {
				return 1
			} else {
				return 0
			}
		}); !found {
			addon = append(addon, &entity.AS{ASN: as, Network: []netip.Prefix{}})
		}
	}

	bgp.AS = append(currentASs, addon...)
	return bgp
}

func ListAvailable() []map[string]any {
	var ls []map[string]any
	if bgpGr := graph.GetBGP(); bgpGr != nil && len(bgpGr.AS) != 0 {
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
