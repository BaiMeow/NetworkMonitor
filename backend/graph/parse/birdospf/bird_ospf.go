package birdospf

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/trace"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/graph/parse/birdospf/parser"
	"github.com/antlr4-go/antlr/v4"
)

//go:generate antlr4 -Dlanguage=Go -visitor -package parser -o parser BirdOSPF.g4

func init() {
	parse.Register("bird-ospf", func(m map[string]any) (parse.Parser[*entity.OSPF], error) {
		return &BirdOSPF{}, nil
	})
}

var _ parse.Parser[*entity.OSPF] = (*BirdOSPF)(nil)

type BirdOSPF struct {
	parse.Base[*entity.OSPF]
}

func (p *BirdOSPF) Parse(ctx context.Context, input any) (*entity.OSPF, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"parse/birdospf/BirdOSPF.Parse",
	)
	defer span.End()

	data, ok := input.([]byte)
	if !ok {
		return nil, fmt.Errorf("input of birdospf parser must be []byte")
	}
	lexer := parser.NewBirdOSPFLexer(antlr.NewIoStream(bytes.NewReader(data)))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	el := &errorListener{}

	streamParser := parser.NewBirdOSPFParser(stream)
	streamParser.AddErrorListener(el)

	tree := streamParser.State()
	visitor := &birdOSPFVisitor{
		graph: new(entity.OSPF),
	}
	state, ok := tree.(*parser.StateContext)
	if !ok {
		return nil, fmt.Errorf("parse as bird ospf state failed")
	}
	visitor.visitState(state)

	if len(el.errs) != 0 {
		err := fmt.Errorf("parse fail")
		for _, e := range el.errs {
			err = fmt.Errorf("%w: %s at line %d, col %d", err, e.msg, e.line, e.col)
		}
		return nil, err
	}
	return visitor.graph, nil
}

type birdOSPFVisitor struct {
	graph *entity.OSPF
}

func (v *birdOSPFVisitor) visitState(ctx *parser.StateContext) {
	for _, area := range ctx.AllArea() {
		v.visitArea(area.(*parser.AreaContext))
	}
}

func (v *birdOSPFVisitor) visitArea(ctx *parser.AreaContext) {
	areaID := ctx.IP().GetText()
	area := v.graph.GetArea(areaID)

	for _, router := range ctx.AllRouter() {
		v.visitRouter(router.(*parser.RouterContext), area)
	}
	for _, network := range ctx.AllNetwork() {
		v.visitNetwork(network.(*parser.NetworkContext), area)
	}
}

func (v *birdOSPFVisitor) visitRouter(ctx *parser.RouterContext, area *entity.Area) {
	if ctx.Distance().Unreachable() != nil {
		return
	}
	routerID := ctx.IP().GetText()
	area.AddRouter(routerID)
	router := area.GetRouter(routerID)

	for _, entry := range ctx.AllRouterEntry() {
		v.visitRouterEntry(entry.(*parser.RouterEntryContext), area, routerID, router)
	}
}

func (v *birdOSPFVisitor) visitRouterEntry(ctx *parser.RouterEntryContext, area *entity.Area, routerID string, router *entity.Router) {
	var cost int
	if i := ctx.INT(0); i != nil {
		c, err := strconv.ParseUint(i.GetText(), 10, 64)
		if err != nil {
			log.Printf("invalid cost %v", i)
		}
		cost = int(c)
	}

	text := ctx.GetText()
	switch {
	case len(text) >= 6 && text[:6] == "router":
		dstRouter := ctx.IP(0).GetText()
		area.AddLink(routerID, dstRouter, cost)
	case len(text) >= 7 && text[:7] == "stubnet":
		prefix := ctx.Prefix().GetText()
		router.AddSubnet(prefix, cost)
	case len(text) >= 8 && text[:8] == "external":
		prefix := ctx.Prefix().GetText()
		router.AddSubnet(prefix, cost)
	case len(text) >= 8 && text[:8] == "xnetwork":
		prefix := ctx.Prefix().GetText()
		router.AddSubnet(prefix, cost)
	case len(text) >= 7 && text[:7] == "xrouter":
		// dstRouter := ctx.IP().GetText()
		// area.AddLink(routerID, dstRouter, cost)
	}
}

func (v *birdOSPFVisitor) visitNetwork(ctx *parser.NetworkContext, area *entity.Area) {
	if ctx.Distance().Unreachable() != nil {
		return
	}
	c, err := strconv.ParseUint(ctx.Distance().INT().GetText(), 10, 64)
	if err != nil {
		log.Printf("invalid distance %v", err)
		return
	}
	allRouterNode := ctx.AllIP()[1:]
	allRouter := make([]string, len(allRouterNode))
	for i, v := range allRouterNode {
		allRouter[i] = v.GetText()
	}
	for _, r1 := range allRouter {
		for _, r2 := range allRouter {
			if r1 != r2 {
				area.AddLink(r1, r2, int(c))
			}
		}
	}
}
