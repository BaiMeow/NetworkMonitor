package graph

import (
	"fmt"
	"github.com/BaiMeow/OSPF-monitor/conf"
	"github.com/BaiMeow/OSPF-monitor/fetch"
	"github.com/BaiMeow/OSPF-monitor/parse"
)

type Probe struct {
	Name   string
	Fetch  fetch.Fetcher
	Parser parse.Parser
	Up     bool
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
		return nil, fmt.Errorf("spawn fetcher fail:%v", err)
	}
	parser, err := parse.Spawn[p.Parse.Type()](p.Parse)
	if err != nil {
		return nil, fmt.Errorf("spawn parser fail:%v", err)
	}
	return &Probe{
		Name:   p.Name,
		Fetch:  fetcher,
		Parser: parser,
		Up:     false,
	}, nil
}

func (p *Probe) DrawAndMerge(drawing *parse.Drawing) (err error) {
	defer func() {
		if err != nil {
			p.Up = false
		} else {
			p.Up = true
		}
	}()
	var data []byte
	data, err = p.Fetch.GetData()
	if err != nil {
		return fmt.Errorf("fetch data from %s fail:%v", p.Name, err)
	}
	p.Parser.Init(data)
	err = p.Parser.ParseAndMerge(drawing)
	if err != nil {
		return fmt.Errorf("parse data from %s fail:%v", p.Name, err)
	}
	return
}
