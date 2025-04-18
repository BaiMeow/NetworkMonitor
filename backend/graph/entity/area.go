package entity

import (
	"slices"
)

type Area struct {
	AreaId string   `json:"area_id,omitempty"`
	Router []Router `json:"router,omitempty"`
	Links  []Link   `json:"links,omitempty"`
}

func (a *Area) AddRouter(RouterId string) {
	for _, router := range a.Router {
		if router.RouterId == RouterId {
			return
		}
	}
	a.Router = append(a.Router, *newRouter(RouterId))
}

func (a *Area) GetRouter(routerId string) *Router {
	for i, router := range a.Router {
		if router.RouterId == routerId {
			return &a.Router[i]
		}
	}
	return nil
}

func (a *Area) AddLink(src, dst string, cost int) {
	a.Links = append(a.Links, newLink(src, dst, cost))
}

func (a *Area) Merge(ar *Area) error {
	if len(a.Router) == 0 {
		a.Router = ar.Router
		a.Links = ar.Links
		return nil
	}

	for _, v := range ar.Router {
		a.AddRouter(v.RouterId)
	}

	var toAdd []Link
	for _, newLink := range ar.Links {
		// ospf always has full graph, merge func is for split ospf
		// if any link existed between the same src and dst, there are all links, so skip.
		if slices.IndexFunc(a.Links, func(link Link) bool {
			return link.Src == newLink.Src && link.Dst == newLink.Dst || link.Dst == newLink.Src && link.Src == newLink.Dst
		}) == -1 {
			toAdd = append(toAdd, newLink)
		}
	}

	for _, newLink := range toAdd {
		a.AddLink(newLink.Src, newLink.Dst, newLink.Cost)
	}
	return nil
}
