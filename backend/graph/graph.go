package graph

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/analysis"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"log"
	"sync"
	"time"
)

var (
	ospf map[uint32]*parse.OSPF
	bgp  *parse.BGP

	bgpBetweenness map[uint32]float64
	bgpCloseness   map[uint32]float64

	currentLock sync.RWMutex
)
var probes []*Probe
var probesLock sync.Mutex

func Init() error {
	for _, probe := range conf.Probes {
		p, err := NewProbe(probe)
		if err != nil {
			return fmt.Errorf("contruct probe fail:%v", err)
		}
		probes = append(probes, p)
	}
	draw()
	ticker := time.NewTicker(conf.Interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				draw()
			}
		}
	}()
	conf.UpdateCallBack = func() {
		probesLock.Lock()
		defer probesLock.Unlock()
		ticker.Reset(conf.Interval)
		var tmp []*Probe
		for _, probe := range conf.Probes {
			p, err := NewProbe(probe)
			if err != nil {
				log.Printf("contruct probe fail:%v\n", err)
			}
			tmp = append(tmp, p)
		}
		probes = tmp
	}
	return nil
}

func draw() {
	var wg sync.WaitGroup
	var drawing parse.Drawing
	drawing.OSPF = make(map[uint32]*parse.OSPF)
	drawing.BGP = &parse.BGP{}

	probesLock.Lock()
	defer probesLock.Unlock()
	for _, p := range probes {
		wg.Add(1)
		p := p
		go func() {
			defer wg.Done()
			alert := time.AfterFunc(conf.ProbeTimeout, func() {
				log.Printf("probe %s timeout but the goroutine is still running, a timeout should be added to the probe.\n", p.Name)
			})
			t := time.Now()
			err := p.DrawAndMerge(&drawing)
			alert.Stop()
			dur := time.Since(t)
			if dur > time.Second*5 {
				log.Printf("probe %s slow draw: %v\n", p.Name, dur)
			}
			if err != nil {
				log.Println(err)
			}
		}()
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

	var (
		tempBetweenness = make(map[uint32]float64)
		tempCloseness   = make(map[uint32]float64)
	)
	drawing.Lock()
	defer drawing.Unlock()
	if conf.Analysis {
		t := time.Now()
		if drawing.BGP != nil {
			g := analysis.ConvertFromBGP(drawing.BGP)
			for _, b := range g.Betweenness() {
				tempBetweenness[b.Node.Tag["asn"].(uint32)] = b.Betweenness
			}
			for _, c := range g.Closeness() {
				tempCloseness[c.Node.Tag["asn"].(uint32)] = c.Closeness
			}
		}
		log.Println("analysis time:", time.Since(t))
	}

	currentLock.Lock()
	defer currentLock.Unlock()
	ospf = drawing.OSPF
	bgp = drawing.BGP
	bgpBetweenness = tempBetweenness
	bgpCloseness = tempCloseness
}
