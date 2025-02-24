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
	disabled bool
	name     string

	data      T
	dataLock  sync.RWMutex
	updatedAt time.Time

	probes     []*Probe[T]
	probesLock sync.Mutex

	Draw         func(ctx context.Context)
	daemonCancel context.CancelFunc
}

func (g *baseGraph[T]) CleanUp() {
	g.disabled = true
	g.daemonCancel()

	for _, p := range g.probes {
		err := p.CleanUp()
		if err != nil {
			log.Printf("stop probe %s fail: %v", p.Name, err)
		}
	}
}

// setProbe must be called under probeLock
func (g *baseGraph[T]) setProbe(probe conf.Probe) error {
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
	g.dataLock.RLock()
	defer g.dataLock.RUnlock()
	return g.data, g.updatedAt
}

func (g *baseGraph[T]) UpdateProbes(confs []conf.Probe) error {
	g.probesLock.Lock()
	defer g.probesLock.Unlock()

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
		err := g.setProbe(c)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if errs != nil {
		return fmt.Errorf("update probes: %v", errs)
	}

	return nil
}

func (g *baseGraph[T]) StartDrawDaemon() {
	ctx, cancel := context.WithCancel(context.Background())
	g.daemonCancel = cancel
	Timer := time.NewTimer(conf.Interval)
	go func() {
		for {
			<-Timer.C
			ctx, cancel := context.WithCancel(ctx)
			alert := time.AfterFunc(conf.ProbeTimeout, func() {
				log.Printf("probe %s timeout, cancelled", g.name)
				cancel()
			})
			g.Draw(ctx)
			t := time.Now()
			alert.Stop()
			dur := time.Since(t)
			if dur > conf.ProbeTimeout/2 {
				log.Printf("probe %s slow draw: %v\n", g.name, dur)
			}
			Timer.Reset(conf.Interval)
		}
	}()
}

type OSPF struct {
	baseGraph[*entity.OSPF]
	asn uint32
}

func newOSPFGraph(asn uint32) *OSPF {
	gr := &OSPF{}
	gr.asn = asn
	gr.baseGraph.Draw = gr.Draw
	gr.baseGraph.name = fmt.Sprintf("AS%d", asn)
	return gr
}

func (o *OSPF) Draw(ctx context.Context) {
	if o.disabled {
		return
	}

	var success int
	data := new(entity.OSPF)

	func() {
		o.probesLock.Lock()
		defer o.probesLock.Unlock()
		for _, p := range o.probes {
			gr, err := p.Draw(ctx)
			if err != nil {
				log.Printf("probe %s error: %v", p.Name, err)
				continue
			}
			success++
			data.Merge(gr)
		}
	}()

	o.dataLock.Lock()
	defer o.dataLock.Unlock()
	if success > 0 {
		o.data = data
		o.updatedAt = time.Now()
	}
}

type BGP struct {
	baseGraph[*entity.BGP]
	betweenness map[uint32]float64
	closeness   map[uint32]float64
}

func newBGPGraph(name string) *BGP {
	gr := &BGP{}
	gr.baseGraph.Draw = gr.Draw
	gr.baseGraph.name = name
	return gr
}

func (b *BGP) Draw(ctx context.Context) {
	if b.disabled {
		return
	}

	var success int
	data := &entity.BGP{
		AS:   make([]*entity.AS, 0),
		Link: make([]entity.ASLink, 0),
	}
	func() {
		b.probesLock.Lock()
		defer b.probesLock.Unlock()
		for _, p := range b.probes {
			e, err := p.Draw(ctx)
			if err != nil {
				log.Printf("probe %s fail: %v", p.Name, err)
				continue
			}
			success++
			data.Merge(e)
		}
	}()

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

	b.dataLock.Lock()
	defer b.dataLock.Unlock()
	if success > 0 {
		b.betweenness = bt
		b.closeness = cl
		b.data = data
		b.updatedAt = time.Now()
	}
}
