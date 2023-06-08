package parse

import "github.com/BaiMeow/NetworkMonitor/conf"

type Router struct {
	RouterId string         `json:"router_id,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

func newRouter(routerId string) *Router {
	return &Router{RouterId: routerId, Metadata: conf.Metas[routerId]}
}
