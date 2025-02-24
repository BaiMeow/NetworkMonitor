package graph

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/service/uptime"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"net/netip"
	"slices"
	"time"
)

func GetOSPF(asn uint32) (*entity.OSPF, time.Time) {
	gr := graph.GetOSPF(asn)
	return gr.GetData()
}

func GetBGP() (*entity.BGP, time.Time) {
	bgp, updatedAt := graph.GetBGP().GetData()
	if bgp == nil {
		return nil, utils.Zero[time.Time]()
	}
	recordASNs, err := uptime.AllASNRecord()
	if err != nil {
		log.Println("read uptime fail:", err)
		return bgp, updatedAt
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
	return bgp, updatedAt
}

func ListAvailable() []map[string]any {
	var ls []map[string]any
	if bgpGr := graph.GetBGP(); bgpGr != nil {
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
