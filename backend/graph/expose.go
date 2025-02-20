package graph

import (
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
)

func GetOSPF(asn uint32) *entity.OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if ospf[asn] == nil {
		return nil
	}
	return ospf[asn].data
}

func GetAllOSPF() map[uint32]*entity.OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	all := make(map[uint32]*entity.OSPF)
	for k, v := range ospf {
		all[k] = v.data
	}
	return all
}

func GetBGP() *entity.BGP {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if len(bgp) == 1 {
		for _, v := range bgp {
			return v.data
		}
	}
	return nil
	// TODO: get bgp by graph name
}

func GetBgpBetweenness() map[uint32]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if len(bgp) == 1 {
		for _, v := range bgp {
			return v.betweenness
		}
	}
	return nil
	// TODO: get bgp by graph name
}

func GetBgpCloseness() map[uint32]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if len(bgp) == 1 {
		for _, v := range bgp {
			return v.closeness
		}
	}
	return nil
	// TODO: get bgp by graph name
}
