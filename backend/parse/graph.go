package parse

type Graph []Area

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

func (g *Graph) Merge(gr *Graph) {
	for _, v := range *gr {
		ar := g.getArea(v.AreaId)
		ar.merge(&v)
	}
}
