package graph

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/analysis"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"log"
	"sync"
	"time"
)

var (
	fullLock sync.RWMutex
	ospf     map[uint32]*OSPF
	bgp      map[string]*BGP
)

type Graph interface {
	AddProbe(probe any)
	Draw()
	CleanUp()
}

type baseGraph[T entity.DrawType] struct {
	lock     sync.RWMutex
	disabled bool

	data      T
	updatedAt time.Time

	probes []*Probe[T]
}

func (g *baseGraph[T]) Draw() {
	panic("implement me")
}

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

func (g *baseGraph[T]) AddProbe(probe any) {
	g.probes = append(g.probes, probe.(*Probe[T]))
}

func (g *baseGraph[T]) GetData() (T, time.Time) {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.data, g.updatedAt
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

	data := new(entity.BGP)
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
