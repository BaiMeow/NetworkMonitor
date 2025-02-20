package entity

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
