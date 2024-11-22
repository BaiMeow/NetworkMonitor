package parse

type Area struct {
	AreaId string   `json:"area_id,omitempty"`
	Router []Router `json:"router,omitempty"`
	Links  []Link   `json:"links,omitempty"`
}

func (a *Area) addRouter(RouterId string) {
	for _, router := range a.Router {
		if router.RouterId == RouterId {
			return
		}
	}
	a.Router = append(a.Router, *newRouter(RouterId))
}

func (a *Area) addLink(src, dst string, cost int) {
	var found bool
	link := newLink(src, dst, cost)
	for _, l := range a.Links {
		if l == link {
			found = true
			break
		}
	}
	if !found {
		a.Links = append(a.Links, link)
	}
}

func (a *Area) merge(ar *Area) {
	if len(a.Router) == 0 {
		a.Router = ar.Router
		a.Links = ar.Links
		return
	}

	for _, v := range ar.Router {
		a.addRouter(v.RouterId)
	}

	// 暂时无法处理两个router间有多条隧道且cost相同的情况
	for _, newLink := range ar.Links {
		a.addLink(newLink.Src, newLink.Dst, newLink.Cost)
	}
}
