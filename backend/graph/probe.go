package graph

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"log"
)

type Drawable interface {
	Draw()
}

type Probe[T entity.DrawType] struct {
	Name   string
	Fetch  fetch.Fetcher
	Parser parse.Parser[T]
}

func NewProbe[T entity.DrawType](p conf.Probe) (any, error) {
	if fetch.Spawn[p.Fetch.Type()] == nil {
		return nil, fmt.Errorf("invalid fetch type:%v", p.Fetch.Type())
	}
	if parse.GetSpawner[T](p.Parse.Type()) == nil {
		return nil, fmt.Errorf("invalid parse type:%v", p.Parse.Type())
	}
	fetcher, err := fetch.Spawn[p.Fetch.Type()](p.Fetch)
	if err != nil {
		return nil, fmt.Errorf("spawn fetcher fail:%v", err)
	}
	parser, err := parse.GetSpawner[T](p.Parse.Type())(p.Parse)
	if err != nil {
		return nil, fmt.Errorf("spawn parser fail:%v", err)
	}
	return &Probe[T]{
		Name:   p.Name,
		Fetch:  fetcher,
		Parser: parser,
	}, nil
}

func (p *Probe[T]) Draw() (T, error) {
	data, err := p.Fetch.GetData()
	if err != nil {
		return nil, fmt.Errorf("fetch data from %s fail:%v", p.Name, err)
	}
	res, err := p.Parser.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse data from %s fail:%v", p.Name, err)
	}
	return res, nil
}

func (p *Probe[T]) CleanUp() error {
	err1 := p.Fetch.CleanUp()
	if err1 != nil {
		log.Printf("stop fetcher: %v", err1)
	}
	err2 := p.Parser.CleanUp()
	if err2 != nil {
		log.Printf("stop parser: %v", err1)
	}
	return nil
}
