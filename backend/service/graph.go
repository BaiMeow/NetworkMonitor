package service

import (
	"github.com/BaiMeow/OSPF-monitor/graph"
	"github.com/BaiMeow/OSPF-monitor/parse"
)

func GetGraph() parse.Graph {
	return graph.Graph
}
