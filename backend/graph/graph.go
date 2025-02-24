package graph

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/analysis"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"log"
	"slices"
	"sync"
	"time"
)

var (
	fullLock sync.RWMutex
	ospf     = make(map[uint32]*OSPF)
	bgp      = make(map[string]*BGP)
)

type baseGraph[T entity.DrawType] struct {
	lock     sync.RWMutex
	disabled bool

	data      T
	updatedAt time.Time

	probes []*Probe[T]

	Draw func(ctx context.Context)
}

//func (g *baseGraph[T]) Draw() {
//	panic("implement me")
//}

func (g *baseGraph[T]) CleanUp() {
	g.disabled = true
	g.lock.Lock()
	g.lock.Unlock()
	for _, p := range g.probes {
		err := p.CleanUp()
		if err != nil {
			log.Printf("stop probe %s fail: %v", p.Name, err)
			return
		}
	}
}

func (g *baseGraph[T]) SetProbe(probe conf.Probe) error {
	idx := slices.IndexFunc(g.probes, func(p *Probe[T]) bool {
		return p.Name == probe.Name
	})
	if idx != -1 && g.probes[idx].conf.Compare(&probe) {
		return nil
	}

	p, err := NewProbe[T](probe)
	if err != nil {
		return fmt.Errorf("contruct probe fail: %v\n", err)
	}

	if idx == -1 {
		g.probes = append(g.probes, p)
	} else {
		go func(oldp *Probe[T]) {
			err := oldp.CleanUp()
			if err != nil {
				log.Printf("clean up %s: %v", oldp.Name, err)
			}
		}(g.probes[idx])
		g.probes[idx] = p
	}
	return nil
}

func (g *baseGraph[T]) GetData() (T, time.Time) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.data, g.updatedAt
}

func (g *baseGraph[T]) UpdateProbes(confs []conf.Probe) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	var errs []error

	g.probes = slices.DeleteFunc(g.probes, func(p *Probe[T]) bool {
		if slices.IndexFunc(confs, func(probe conf.Probe) bool {
			return probe.Name == p.Name
		}) == -1 {
			go func(oldp *Probe[T]) {
				err := oldp.CleanUp()
				if err != nil {
					log.Printf("clean up %s: %v", oldp.Name, err)
				}
			}(p)
			return true
		}
		return false
	})

	for _, c := range confs {
		err := g.SetProbe(c)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if errs != nil {
		return fmt.Errorf("update probes: %v", errs)
	}

	return nil
}

type OSPF struct {
	baseGraph[*entity.OSPF]
	asn uint32
}

func (o *OSPF) Draw() {
	if o.disabled {
		return
	}
	data := new(entity.OSPF)
	for _, p := range o.probes {
		gr, err := p.Draw()
		if err != nil {
			log.Printf("probe %s error: %v", p.Name, err)
			continue
		}
		data.Merge(gr)
	}
	o.lock.Lock()
	defer o.lock.Unlock()
	o.data = data
	o.updatedAt = time.Now()
}

type BGP struct {
	baseGraph[*entity.BGP]
	name        string
	betweenness map[uint32]float64
	closeness   map[uint32]float64
}

func (b *BGP) Draw() {
	if b.disabled {
		return
	}

	data := &entity.BGP{
		AS:   make([]*entity.AS, 0),
		Link: make([]entity.ASLink, 0),
	}
	for _, p := range b.probes {
		e, err := p.Draw()
		if err != nil {
			log.Printf("probe %s fail: %v", p.Name, err)
			continue
		}
		data.Merge(e)
	}

	bt := make(map[uint32]float64)
	cl := make(map[uint32]float64)
	if conf.Analysis {
		t := time.Now()
		g := analysis.ConvertFromBGP(data)

		for _, b := range g.Betweenness() {
			bt[b.Node.Tag["asn"].(uint32)] = b.Betweenness
		}
		for _, c := range g.Closeness() {
			cl[c.Node.Tag["asn"].(uint32)] = c.Closeness
		}
		log.Println("analysis time:", time.Since(t))
	}

	b.lock.Lock()
	defer b.lock.Unlock()
	b.betweenness = bt
	b.closeness = cl
	b.data = data
	b.updatedAt = time.Now()
}
