package parse

type Router struct {
	RouterId string `json:"router_id,omitempty"`
}

func newRouter(routerId string) *Router {
	return &Router{RouterId: routerId}
}
