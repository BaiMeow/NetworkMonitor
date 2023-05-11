package graph

import (
	"fmt"
	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/parse"
	"log"
	"time"
)

var Graph parse.Graph

var probes []*Probe

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
	return nil
}

func drawGraph() {
	var gh parse.Graph
	for _, p := range probes {
		graph, err := p.GetGraph()
		if err != nil {
			log.Println(err)
			continue
		}
		gh.Merge(&graph)
	}
	Graph = gh
}
