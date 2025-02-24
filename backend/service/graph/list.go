package graph

import "github.com/BaiMeow/NetworkMonitor/graph"

func ListAvailable() []map[string]any {
	var ls []map[string]any
	for name := range graph.GetAllBGP() {
		ls = append(ls, map[string]any{
			"type": "bgp",
			"name": name,
		})
	}

	for asn := range graph.GetAllOSPF() {
		ls = append(ls, map[string]any{
			"type": "ospf",
			"asn":  asn,
		})
	}

	return ls
}
