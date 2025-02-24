package graph

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"maps"
	"strconv"
	"sync"
	"time"
)

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
		return ospfNew[k] == nil
	})
	maps.DeleteFunc(bgp, func(k string, v *BGP) bool {
		return bgpNew[k] == nil
	})
	for k, v := range ospfNew {
		if ospf[k] == nil {
			ospf[k] = &OSPF{asn: k}
		}
		err := ospf[k].UpdateProbes(v)
		if err != nil {
			log.Printf("config ospf graph %d fail: %v", k, err)
			continue
		}
	}
	for k, v := range bgpNew {
		if bgp[k] == nil {
			bgp[k] = &BGP{name: k}
		}
		err := bgp[k].UpdateProbes(v)
		if err != nil {
			log.Printf("config bgp graph %s fail: %v", k, err)
			continue
		}
	}
	log.Println("update probes done")
}

func Init() error {
	// update graph
	patchGraphList()
	draw()

	ticker := time.NewTicker(conf.Interval)
	go func() {
		for {
			<-ticker.C
			draw()
		}
	}()

	conf.UpdateCallBack = func() {
		ticker.Reset(conf.Interval)
		for _, gr := range ospf {
			go gr.CleanUp()
		}
		for _, gr := range bgp {
			go gr.CleanUp()
		}
		patchGraphList()
	}
	return nil
}

func draw() {
	var wg sync.WaitGroup
	fullLock.RLock()
	defer fullLock.RUnlock()
	drawSingle := func(name string, drawable Drawable) {
		defer wg.Done()
		alert := time.AfterFunc(conf.ProbeTimeout, func() {
			log.Printf("probe %s timeout but the goroutine is still running, a timeout should be added to the probe.\n", name)
		})
		drawable.Draw()
		t := time.Now()
		alert.Stop()
		dur := time.Since(t)
		if dur > time.Second*5 {
			log.Printf("probe %s slow draw: %v\n", name, dur)
		}
	}
	for name, v := range bgp {
		wg.Add(1)
		go drawSingle(name, v)
	}
	for name, v := range ospf {
		wg.Add(1)
		go drawSingle(strconv.FormatUint(uint64(name), 10), v)
	}

	select {
	case <-time.After(conf.ProbeTimeout):
	case <-func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		return done
	}():
	}
}
