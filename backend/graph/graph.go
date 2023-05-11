package graph

import (
	"fmt"
	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/parse"
	"log"
	"sync"
	"time"
)

var Graph parse.Graph

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
	drawGraph()
	ticker := time.NewTicker(time.Second * time.Duration(conf.Interval))
	go func() {
		for {
			select {
			case <-ticker.C:
				drawGraph()
			}
		}
	}()
	conf.UpdateCallBack = func() {
		probesLock.Lock()
		defer probesLock.Unlock()
		ticker.Reset(time.Duration(conf.Interval) * time.Second)
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

func drawGraph() {
	var gh parse.Graph
	var lock sync.Mutex
	var wg sync.WaitGroup

	probesLock.Lock()
	defer probesLock.Unlock()
	for _, p := range probes {
		wg.Add(1)
		p := p
		go func() {
			defer wg.Done()
			graph, err := p.GetGraph()
			if err != nil {
				log.Println(err)
			}
			lock.Lock()
			defer lock.Unlock()
			gh.Merge(&graph)
		}()
	}

	wg.Wait()
	Graph = gh
}
