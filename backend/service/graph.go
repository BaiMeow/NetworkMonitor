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

func ListAvailable() []map[string]any {
	var ls []map[string]any
	if len(graph.BGP.AS) != 0 {
		ls = append(ls, map[string]any{
			"type": "bgp",
		})
	}

	for asn := range graph.OSPF {
		ls = append(ls, map[string]any{
			"type": "ospf",
			"asn":  asn,
		})
	}

	return ls
}
