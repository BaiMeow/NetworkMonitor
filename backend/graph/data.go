package graph

import "github.com/BaiMeow/NetworkMonitor/graph/parse"

func GetOSPF(asn uint32) *parse.OSPF {
	currentLock.RLock()
	defer currentLock.RUnlock()
	return ospf[asn]
}

func GetAllOSPF() map[uint32]*parse.OSPF {
	currentLock.RLock()
	defer currentLock.RUnlock()
	return ospf
}

func GetBGP() *parse.BGP {
	currentLock.RLock()
	defer currentLock.RUnlock()
	bgp := *bgp
	return &bgp
}

func GetBgpBetweenness() map[uint32]float64 {
	currentLock.RLock()
	defer currentLock.RUnlock()
	return bgpBetweenness
}

func GetBgpCloseness() map[uint32]float64 {
	currentLock.RLock()
	defer currentLock.RUnlock()
	return bgpCloseness
}
