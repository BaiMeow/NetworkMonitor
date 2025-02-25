package graph

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/analysis"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"log"
	"slices"
	"strconv"
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
	if idx != -1 && g.probes[idx].conf.Equal(&probe) {
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
	waitFirstLoad := make(chan struct{})
	go func() {
		once := sync.OnceFunc(func() {
			close(waitFirstLoad)
		})
		for {
			ctx, cancel := context.WithCancel(ctx)
			t := time.Now()
			alert := time.AfterFunc(conf.ProbeTimeout, func() {
				log.Printf("graph %s timeout, cancelled", g.name)
				cancel()
			})
			g.Draw(ctx)
			alert.Stop()
			dur := time.Since(t)
			if dur > conf.ProbeTimeout/2 {
				log.Printf("graph %s slow draw: %v\n", g.name, dur)
			}
			once()
			Timer.Reset(conf.Interval)
			<-Timer.C
		}
	}()
	<-waitFirstLoad
}

type OSPF struct {
	baseGraph[*entity.OSPF]
	asn uint32
}

func newOSPFGraph(asn uint32) *OSPF {
	gr := &OSPF{
		baseGraph: baseGraph[*entity.OSPF]{
			name: strconv.FormatUint(uint64(asn), 10),
			data: &entity.OSPF{},
		},
		asn: asn,
	}
	gr.asn = asn
	gr.baseGraph.Draw = gr.Draw
	gr.baseGraph.name = strconv.FormatUint(uint64(asn), 10)
	return gr
}

func (o *OSPF) Draw(ctx context.Context) {
	if o.disabled {
		return
	}

	var success int
	data := new(entity.OSPF)
	var drawLock sync.Mutex

	func() {
		o.probesLock.Lock()
		defer o.probesLock.Unlock()
		var wg sync.WaitGroup
		for _, p := range o.probes {
			wg.Add(1)
			go func() {
				defer wg.Done()
				gr, err := p.Draw(ctx)
				if err != nil {
					log.Printf("probe %s error: %v", p.Name, err)
					return
				}
				success++
				drawLock.Lock()
				defer drawLock.Unlock()
				data.Merge(gr)
			}()
		}
		wg.Wait()
	}()

	o.dataLock.Lock()
	defer o.dataLock.Unlock()
	if success > 0 {
		equal := o.data.Equal(data)
		o.data = data
		o.updatedAt = time.Now()
		if !equal {
			notifyEventUpdate("ospf", o.name)
		}
	}
}

type BGP struct {
	baseGraph[*entity.BGP]
	betweenness map[uint32]float64
	closeness   map[uint32]float64
}

func newBGPGraph(name string) *BGP {
	gr := &BGP{
		baseGraph: baseGraph[*entity.BGP]{
			name: name,
			data: &entity.BGP{
				AS:   make([]*entity.AS, 0),
				Link: make([]entity.ASLink, 0),
			},
		},
	}
	gr.baseGraph.Draw = gr.Draw
	return gr
}

func (b *BGP) Draw(ctx context.Context) {
	if b.disabled {
		return
	}

	var success int
	var drawLock sync.Mutex
	data := &entity.BGP{
		AS:   make([]*entity.AS, 0),
		Link: make([]entity.ASLink, 0),
	}

	func() {
		b.probesLock.Lock()
		defer b.probesLock.Unlock()
		var wg sync.WaitGroup
		for _, p := range b.probes {
			wg.Add(1)
			go func() {
				defer wg.Done()
				e, err := p.Draw(ctx)
				if err != nil {
					log.Printf("probe %s fail: %v", p.Name, err)
					return
				}
				success++
				drawLock.Lock()
				defer drawLock.Unlock()
				data.Merge(e)
			}()
		}
		wg.Wait()
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
		equal := b.data.Equal(data)
		b.data = data
		b.updatedAt = time.Now()
		if !equal {
			notifyEventUpdate("bgp", b.name)
		}
	}
}
