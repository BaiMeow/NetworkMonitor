package birdospf

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/graph/parse/birdospf/parser"
	"github.com/antlr4-go/antlr/v4"
)

//go:generate antlr4 -Dlanguage=Go -visitor -package parser -o parser BirdOSPF.g4

func init() {
	parse.Register("bird-ospf", func(m map[string]any) (parse.Parser, error) {
		asn, ok := m["asn"].(int)
		if !ok {
			return nil, fmt.Errorf("asn is not int")
		}
		return &BirdOSPF{
			asn: uint32(asn),
		}, nil
	})
}

var _ parse.Parser = (*BirdOSPF)(nil)

type BirdOSPF struct {
	lexer  *parser.BirdOSPFLexer
	parser *parser.BirdOSPFParser
	errL   *errorListener
	asn    uint32
}

func (p *BirdOSPF) Init(input []byte) {
	lexer := parser.NewBirdOSPFLexer(antlr.NewIoStream(bytes.NewReader(input)))
	stream := antlr.NewCommonTokenStream(lexer, 0)
	el := &errorListener{}

	p.lexer = lexer
	p.parser = parser.NewBirdOSPFParser(stream)
	p.errL = el
	p.parser.AddErrorListener(el)
}

func (p *BirdOSPF) ParseAndMerge(drawing *parse.Drawing) (err error) {
	tree := p.parser.State()
	visitor := &birdOSPFVisitor{
		graph: new(parse.OSPF),
	}
	state, ok := tree.(*parser.StateContext)
	if !ok {
		return fmt.Errorf("parse as bird ospf state failed")
	}
	visitor.visitState(state)
	drawing.Lock()
	defer drawing.Unlock()
	drawing.OSPF[p.asn] = visitor.graph

	if len(p.errL.errs) != 0 {
		err := fmt.Errorf("parse fail")
		for _, e := range p.errL.errs {
			err = fmt.Errorf("%w: %s at line %d, col %d", err, e.msg, e.line, e.col)
		}
		return err
	}
	return nil
}

type birdOSPFVisitor struct {
	graph *parse.OSPF
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
}

func (v *birdOSPFVisitor) visitRouter(ctx *parser.RouterContext, area *parse.Area) {
	routerID := ctx.IP().GetText()
	area.AddRouter(routerID)
	router := area.GetRouter(routerID)

	for _, entry := range ctx.AllRouterEntry() {
		v.visitRouterEntry(entry.(*parser.RouterEntryContext), area, routerID, router)
	}
}

func (v *birdOSPFVisitor) visitRouterEntry(ctx *parser.RouterEntryContext, area *parse.Area, routerID string, router *parse.Router) {
	var cost int
	if i := ctx.INT(); i != nil {
		c, err := strconv.ParseUint(i.GetText(), 10, 64)
		if err != nil {
			fmt.Printf("invalid cost %v", i)
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
	case len(text) >= 9 && text[:9] == "xnetwork":
		prefix := ctx.Prefix().GetText()
		router.AddSubnet(prefix, cost)
	case len(text) >= 7 && text[:7] == "xrouter":
		// dstRouter := ctx.IP().GetText()
		// area.AddLink(routerID, dstRouter, cost)
	}
}
