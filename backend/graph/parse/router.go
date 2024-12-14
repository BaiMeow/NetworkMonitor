package parse

import (
	"net/netip"
	"slices"
)

type Router struct {
	RouterId string         `json:"router_id,omitempty"`
	Subnet   []netip.Prefix `json:"subnet,omitempty"`
}

func newRouter(routerId string) *Router {
	return &Router{RouterId: routerId}
}
func (r *Router) AddSubnet(network string, metric int) {
	prefix, _ := netip.ParsePrefix(network)
	if slices.Index(r.Subnet, prefix) == -1 {
		r.Subnet = append(r.Subnet, prefix)
	}
}
