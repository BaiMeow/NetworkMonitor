package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	Register("bird-ospf", func(m map[string]any) (Parser, error) {
		asn, ok := m["asn"].(int)
		if !ok {
			return nil, fmt.Errorf("asn is not int")
		}
		return &BirdOSPF{
			asn: uint32(asn),
		}, nil
	})
}

type BirdOSPF struct {
	s     *bufio.Scanner
	graph OSPF
	// ctx
	asn    uint32
	area   string
	router string
}

func (p *BirdOSPF) Init(input []byte) {
	p.s = bufio.NewScanner(bytes.NewReader(input))
	p.graph = nil
	p.area = ""
	p.router = ""
}

func (p *BirdOSPF) ParseAndMerge(drawing *Drawing) (err error) {
	defer func() {
		if err != nil {
			return
		}
		drawing.Lock()
		defer drawing.Unlock()
		if ospf, ok := drawing.OSPF[p.asn]; !ok {
			drawing.OSPF[p.asn] = &p.graph
		} else {
			ospf.Merge(&p.graph)
		}
	}()
	for p.s.Scan() {
		fields := strings.Split(strings.TrimSpace(p.s.Text()), " ")
		switch fields[0] {
		case "BIRD":
			// skip next line
			p.s.Scan()
		case "area":
			err = p.parseArea(fields)
		case "router":
			err = p.parseRouter(fields)
		case "other":
			// skip other areas
			return nil
		case "unreachable":
			err = p.parseUnreachable(fields)
		case "":
			p.left()
		}
		if err != nil {
			return
		}
	}
	return nil
}

func (p *BirdOSPF) parseArea(fields []string) error {
	if len(fields) != 2 {
		return fmt.Errorf("invalid bird format:%v", fields)
	}
	p.area = fields[1]
	p.s.Scan()
	return nil
}

func (p *BirdOSPF) parseRouter(fields []string) error {
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

func (p *BirdOSPF) skip(words int) bool {
	for ; words > 0; words-- {
		if !p.s.Scan() {
			return false
		}
	}
	return true
}

func (p *BirdOSPF) left() {
	if p.router != "" {
		p.router = ""
		return
	}
}

func (p *BirdOSPF) parseUnreachable(fields []string) error {
	if len(fields) != 1 {
		return fmt.Errorf("invalid bird format:%v", fields)
	}
	p.router = ""
	for {
		if p.s.Text() == "" || !p.s.Scan() {
			break
		}
	}
	return nil
}
