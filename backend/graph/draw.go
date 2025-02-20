package graph

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"strconv"
	"sync"
	"time"
)

func newAndAddProbe[T entity.DrawType](graph Graph, probe conf.Probe) error {
	p, err := NewProbe[T](probe)
	if err != nil {
		return fmt.Errorf("contruct probe fail: %v\n", err)
	}
	graph.AddProbe(p)
	return nil
}

func recreateGraph() {
	ospfNew := make(map[uint32]*OSPF)
	bgpNew := make(map[string]*BGP)

	for _, probe := range conf.Probes {
		switch probe.Draw.Type() {
		case "ospf":
			asn, err := utils.MustASN(probe.Draw["asn"])
			if err != nil {
				log.Printf("parse drawer failed: %v\n", err)
				continue
			}
			if _, ok := ospfNew[asn]; !ok {
				ospfNew[asn] = &OSPF{asn: asn}
			}
			if err := newAndAddProbe[*entity.OSPF](ospfNew[asn], probe); err != nil {
				log.Printf("add probe %s: %v", probe.Name, err)
				continue
			}
		case "bgp":
			name, ok := probe.Draw["name"].(string)
			if !ok {
				log.Printf("bgp graph name field not found")
				continue
			}
			if _, ok := bgpNew[name]; !ok {
				bgpNew[name] = &BGP{name: name}
			}
			if err := newAndAddProbe[*entity.BGP](bgpNew[name], probe); err != nil {
				log.Printf("add probe %s: %v", probe.Name, err)
				continue
			}
		default:
			log.Printf("unknown draw type %s", probe.Draw.Type())
		}
	}

	fullLock.Lock()
	bgp = bgpNew
	ospf = ospfNew
	fullLock.Unlock()
}

func Init() error {
	// update graph
	recreateGraph()
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
		recreateGraph()
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
