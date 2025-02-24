package graph

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"maps"
)

func Init() error {
	// update graph
	patchGraphList()
	conf.UpdateCallBack = patchGraphList
	return nil
}

func patchGraphList() {
	ospfNew := make(map[uint32][]conf.Probe)
	bgpNew := make(map[string][]conf.Probe)

	for _, probe := range conf.Probes {
		switch probe.Draw.Type() {
		case "ospf":
			asn, err := utils.MustASN(probe.Draw["asn"])
			if err != nil {
				log.Printf("parse drawer failed: %v\n", err)
				continue
			}
			ospfNew[asn] = append(ospfNew[asn], probe)
		case "bgp":
			name, ok := probe.Draw["name"].(string)
			if !ok {
				log.Printf("bgp graph name field not found")
				continue
			}
			bgpNew[name] = append(bgpNew[name], probe)
		default:
			log.Printf("unknown draw type %s", probe.Draw.Type())
		}
	}

	log.Println("start updating probes, service paused")
	fullLock.Lock()
	defer fullLock.Unlock()
	maps.DeleteFunc(ospf, func(k uint32, v *OSPF) bool {
		if ospfNew[k] == nil {
			go v.CleanUp()
			return true
		}
		return false
	})
	maps.DeleteFunc(bgp, func(k string, v *BGP) bool {
		if bgpNew[k] == nil {
			go v.CleanUp()
			return true
		}
		return false
	})
	for k, v := range ospfNew {
		isNewProbe := ospf[k] == nil
		if isNewProbe {
			ospf[k] = newOSPFGraph(k)
		}
		err := ospf[k].UpdateProbes(v)
		if err != nil {
			log.Printf("config ospf graph %d fail: %v", k, err)
			continue
		}
		if isNewProbe {
			go ospf[k].StartDrawDaemon()
		}
	}
	for k, v := range bgpNew {
		isNewProbe := bgp[k] == nil
		if isNewProbe {
			bgp[k] = newBGPGraph(k)
		}
		err := bgp[k].UpdateProbes(v)
		if err != nil {
			log.Printf("config bgp graph %s fail: %v", k, err)
			continue
		}
		if isNewProbe {
			go bgp[k].StartDrawDaemon()
		}
	}
	log.Println("update probes done")
}
