package conf

import "time"

type Probe struct {
	Name  string
	Parse Parser
	Fetch Fetcher
}

type Parser map[string]any

func (p Parser) Type() string {
	return p["type"].(string)
}

type Fetcher map[string]any

func (f Fetcher) Type() string {
	return f["type"].(string)
}

type UptimeCfg struct {
	StoreDuration time.Duration
	Interval      time.Duration
}
