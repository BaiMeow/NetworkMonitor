package service

import (
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/parse"
)

func GetOSPF(asn uint32) *parse.OSPF {
	return graph.OSPF[asn]
}

func GetBGP() *parse.BGP {
	return &graph.BGP
}
