package parse

type Area struct {
	AreaId string   `json:"area_id,omitempty"`
	Router []Router `json:"router,omitempty"`
	Links  []Link   `json:"links,omitempty"`
}

func (a *Area) FromANTLRContext(ctx interface{}) error {
	// 这里将添加从ANTLR Context解析数据的逻辑
	return nil
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

func (a *Area) AddLink(src, dst string, cost int) error {
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
	return nil
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

	for _, newLink := range ar.Links {
		if err := a.AddLink(newLink.Src, newLink.Dst, newLink.Cost); err != nil {
			return err
		}
	}
	return nil
}
