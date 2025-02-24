package entity

import (
	"maps"
	"net/netip"
	"slices"
)

type OSPF []Area

func (g *OSPF) GetArea(areaId string) *Area {
	for i, area := range *g {
		if area.AreaId == areaId {
			return &(*g)[i]
		}
	}
	*g = append(*g, Area{AreaId: areaId})
	return &(*g)[len(*g)-1]
}

func (g *OSPF) AddLink(area, src, dst string, cost int) {
	g.GetArea(area).AddLink(src, dst, cost)
}

func (g *OSPF) Merge(gr *OSPF) {
	for _, v := range *gr {
		ar := g.GetArea(v.AreaId)
		ar.Merge(&v)
	}
}

func (g *OSPF) Equal(g2 *OSPF) bool {
	if len(*g) != len(*g2) {
		return false
	}

	gmap := maps.Collect(func(yield func(string, Area) bool) {
		slices.All(*g)(func(_ int, area Area) bool {
			return yield(area.AreaId, area)
		})
	})

	for _, v := range *g2 {
		peer, ok := gmap[v.AreaId]
		if !ok {
			return false
		}

		if len(v.Links) != len(peer.Links) {
			return false
		}
		linkMap := make(map[Link]int)
		for _, vlink := range v.Links {
			linkMap[vlink]++
		}
		for _, peerLink := range peer.Links {
			linkMap[peerLink]--
		}
		for _, v := range linkMap {
			if v != 0 {
				return false
			}
		}

		vRouters := maps.Collect(func(yield func(string, Router) bool) {
			slices.All(v.Router)(func(_ int, router Router) bool {
				return yield(router.RouterId, router)
			})
		})

		for _, peerRouter := range peer.Router {
			vRouter, ok := vRouters[peerRouter.RouterId]
			if !ok {
				return false
			}
			if slices.IndexFunc(vRouter.Subnet, func(prefix netip.Prefix) bool {
				return slices.IndexFunc(peerRouter.Subnet, func(prefix2 netip.Prefix) bool {
					return prefix2 == prefix
				}) == -1
			}) != -1 {
				return false
			}
		}
	}
	return true
}
