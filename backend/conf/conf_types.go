package conf

import (
	"reflect"
	"time"
)

type Probe struct {
	Name  string
	Parse Parser
	Fetch Fetcher
	Draw  Drawer
}

func (p *Probe) Equal(p2 *Probe) bool {
	return reflect.DeepEqual(p.Name, p2.Name)
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
