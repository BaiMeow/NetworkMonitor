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
		data, err := p.Fetch.GetData()
		if err != nil {
			log.Printf("fetch fail: %v", err)
			continue
		}
		p.Parser.Init(data)
		graph, err := p.Parser.Parse()
		if err != nil {
			log.Printf("parse fail:%v", err)
			continue
		}
		gh.Merge(&graph)
	}
	Graph = gh
}
