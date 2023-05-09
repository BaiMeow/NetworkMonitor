package parse

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	Register("bird", func(m map[string]any) (Parser, error) {
		return &BirdParser{}, nil
	})
}

type BirdParser struct {
	s     *bufio.Scanner
	graph Graph
	// ctx
	area   string
	router string
}

func (p *BirdParser) Init(input string) {
	p.s = bufio.NewScanner(strings.NewReader(input))
	p.graph = nil
	p.area = ""
	p.router = ""
}

func (p *BirdParser) Parse() (Graph, error) {
	for p.s.Scan() {
		fields := strings.Split(strings.TrimSpace(p.s.Text()), " ")
		var err error
		switch fields[0] {
		case "BIRD":
			// skip next line
			p.s.Scan()
		case "area":
			err = p.parseArea(fields)
		case "router":
			err = p.parseRouter(fields)
		case "":
			p.left()
		}
		if err != nil {
			return nil, err
		}
	}
	return p.graph, nil
}

func (p *BirdParser) parseArea(fields []string) error {
	if len(fields) != 2 {
		return fmt.Errorf("invalid bird format:%v", fields)
	}
	p.area = fields[1]
	p.s.Scan()
	return nil
}

func (p *BirdParser) parseRouter(fields []string) error {
	if len(fields) == 2 && p.router == "" && p.area != "" {
		p.router = fields[1]
		p.graph.getArea(p.area).addRouter(p.router)
		return nil
	}

	if len(fields) == 4 && p.router != "" && p.area != "" {
		router := fields[1]
		cost, err := strconv.Atoi(fields[3])
		if err != nil {
			return fmt.Errorf("invalid bird format:%v", err)
		}
		p.graph.addLink(p.area, p.router, router, cost)
		return nil
	}

	return fmt.Errorf("invalid bird format:%v", fields)
}

func (p *BirdParser) skip(words int) {
	for ; words > 0; words-- {
		p.s.Scan()
	}
}

func (p *BirdParser) left() {
	if p.router != "" {
		p.router = ""
		return
	}
}
