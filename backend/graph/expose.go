package graph

func GetOSPF(asn uint32) *OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if ospf[asn] == nil {
		return nil
	}
	return ospf[asn]
}

func GetAllOSPF() map[uint32]*OSPF {
	fullLock.RLock()
	defer fullLock.RUnlock()
	all := make(map[uint32]*OSPF)
	for k, v := range ospf {
		all[k] = v
	}
	return all
}

func GetBGP() *BGP {
	fullLock.RLock()
	defer fullLock.RUnlock()
	if len(bgp) == 1 {
		for _, v := range bgp {
			return v
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
