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

func GetBGP(name string) (*entity.BGP, time.Time) {
	gr := graph.GetBGP(name)
	if gr == nil {
		return nil, utils.Zero[time.Time]()
	}
	bgp, updatedAt := gr.GetData()
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
