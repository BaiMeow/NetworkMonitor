package conf

import (
	"maps"
	"time"
)

type Probe struct {
	Name  string
	Parse Parser
	Fetch Fetcher
	Draw  Drawer
}

func (p *Probe) Equal(p2 *Probe) bool {
	if p == nil && p2 != nil || p != nil && p2 == nil {
		return false
	}
	if p == nil {
		return true
	}
	return p.Name == p2.Name &&
		maps.Equal(p.Parse, p2.Parse) &&
		maps.Equal(p.Fetch, p2.Fetch) &&
		maps.Equal(p.Draw, p2.Draw)
}

type Parser map[string]any

func (p Parser) Type() string {
	return p["type"].(string)
}

type Fetcher map[string]any

func (f Fetcher) Type() string {
	return f["type"].(string)
}

type Drawer map[string]any

func (f Drawer) Type() string {
	return f["type"].(string)
}

type UptimeCfg struct {
	StoreDuration time.Duration
	Interval      time.Duration
}

type Tracer struct {
	Endpoint    string
	ServiceName string
}

type IOT struct {
	Port    int
	Enabled bool
}
