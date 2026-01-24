package graph

import (
	"maps"
)

func GetOSPF(asn uint32) *OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	return ospf[asn]
}

func GetAllOSPF() map[uint32]*OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	return maps.Clone(ospf)
}

func GetBGP(name string) *BGP {
	fullLock.RLock()
	defer fullLock.RUnlock()
	return bgp[name]
}

func GetAllBGP() map[string]*BGP {
	fullLock.RLock()
	defer fullLock.RUnlock()
	return maps.Clone(bgp)
}

func GetBgpBetweenness(name string) map[uint32]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := bgp[name]
	if gr != nil {
		return gr.betweenness
	}
	return nil
}

func GetBgpCloseness(name string) map[uint32]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := bgp[name]
	if gr != nil {
		return gr.closeness
	}
	return nil
}

func GetBgpPathBetweenness(name string) []PathBetweenness {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := bgp[name]
	if gr != nil {
		return gr.pathBetweenness
	}
	return nil
}

func GetOSPFBetweenness(asn uint32) map[string]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := ospf[asn]
	if gr != nil {
		return gr.betweenness
	}
	return nil
}

func GetOSPFCloseness(asn uint32) map[string]float64 {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := ospf[asn]
	if gr != nil {
		return gr.closeness
	}
	return nil
}

func GetOSPFPathBetweenness(asn uint32) []PathBetweenness {
	fullLock.RLock()
	defer fullLock.RUnlock()
	gr := ospf[asn]
	if gr != nil {
		return gr.pathBetweenness
	}
	return nil
}
