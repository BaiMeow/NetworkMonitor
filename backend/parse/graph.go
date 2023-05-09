package parse

type Graph []Area

type Area struct {
	AreaId string   `json:"area_id,omitempty"`
	Router []Router `json:"router,omitempty"`
	Links  []Link   `json:"links,omitempty"`
}

type Router struct {
	RouterId string `json:"router_id,omitempty"`
}

func (a *Area) addRouter(RouterId string) {
	for _, router := range a.Router {
		if router.RouterId == RouterId {
			return
		}
	}
	a.Router = append(a.Router, Router{RouterId: RouterId})
}

func (g *Graph) getArea(areaId string) *Area {
	for i, area := range *g {
		if area.AreaId == areaId {
			return &(*g)[i]
		}
	}
	*g = append(*g, Area{AreaId: areaId})
	return &(*g)[len(*g)-1]
}

func (g *Graph) addLink(area, src, dst string, cost int) {
	g.getArea(area).addLink(src, dst, cost)
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

func (g *Graph) Merge(gr *Graph) {
	for _, v := range *gr {
		ar := g.getArea(v.AreaId)
		ar.merge(&v)
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
