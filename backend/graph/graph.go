package graph

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	parse2 "github.com/BaiMeow/NetworkMonitor/graph/parse"
	"log"
	"sync"
	"time"
)

var OSPF map[uint32]*parse2.OSPF
var BGP *parse2.BGP

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
	var drawing parse2.Drawing
	drawing.OSPF = make(map[uint32]*parse2.OSPF)
	drawing.BGP = &parse2.BGP{}

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

	drawing.Lock()
	defer drawing.Unlock()
	OSPF = drawing.OSPF
	BGP = drawing.BGP
}
