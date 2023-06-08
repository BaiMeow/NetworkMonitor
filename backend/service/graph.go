package service

import (
	"github.com/BaiMeow/OSPF-monitor/graph"
	"github.com/BaiMeow/OSPF-monitor/parse"
)

func GetOSPF(asn uint32) *parse.OSPF {
	return graph.OSPF[asn]
}

func GetBGP() *parse.BGP {
	return &graph.BGP
}
