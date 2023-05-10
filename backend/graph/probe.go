package graph

import (
	"fmt"
	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/fetch"
	"github.com/BaiMeow/OSPF-monitor/parse"
)

type Probe struct {
	Fetch  fetch.Fetcher
	Parser parse.Parser
}

func NewProbe(p conf.Probe) (*Probe, error) {
	if fetch.Spawn[p.Fetch.Type()] == nil {
		return nil, fmt.Errorf("invalid fetch type:%v", p.Fetch.Type())
	}
	if parse.Spawn[p.Parse.Type()] == nil {
		return nil, fmt.Errorf("invalid parse type:%v", p.Parse.Type())
	}
	fetcher, err := fetch.Spawn[p.Fetch.Type()](p.Fetch)
	if err != nil {
		return nil, fmt.Errorf("spawn fetcher fail:%v", fetcher)
	}
	parser, err := parse.Spawn[p.Parse.Type()](p.Parse)
	if err != nil {
		return nil, fmt.Errorf("spawn parser fail:%v", parser)
	}
	return &Probe{
		Fetch:  fetcher,
		Parser: parser,
	}, nil
}
