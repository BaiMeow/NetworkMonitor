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

var _ Parser = (*BirdOSPF)(nil)

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
		case "network":
			err = p.parseNetwork(fields)
		case "":
			p.leftRouter()
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

func (p *BirdOSPF) skip(lines int) bool {
	for ; lines > 0; lines-- {
		if !p.s.Scan() {
			return false
		}
	}
	return true
}

func (p *BirdOSPF) leftRouter() {
	p.router = ""
}

func (p *BirdOSPF) parseUnreachable(fields []string) error {
	if len(fields) != 1 {
		return fmt.Errorf("invalid bird format:%v", fields)
	}
	for {
		if p.s.Text() == "" || !p.s.Scan() {
			break
		}
	}
	p.leftRouter()
	return nil
}

func (p *BirdOSPF) parseNetwork(fields []string) error {
	if len(fields) == 4 && p.router != "" && p.area != "" {
		return nil
	}

	if len(fields) != 2 {
		return fmt.Errorf("invalid bird format:%v", fields)
	}

	// read and skip dr
	p.s.Scan()

	// read distance
	p.s.Scan()
	fields = strings.Split(strings.TrimSpace(p.s.Text()), " ")
	if len(fields) != 2 {
		return fmt.Errorf("invalid bird format:%s", p.s.Text())
	}

	distance, err := strconv.Atoi(fields[1])
	if err != nil {
		return fmt.Errorf("invalid bird format:%s", p.s.Text())
	}

	var routers []string

	for {
		if p.s.Text() == "" || !p.s.Scan() {
			break
		}
		fields = strings.Split(strings.TrimSpace(p.s.Text()), " ")
		if len(fields) == 2 && fields[0] == "router" {
			routers = append(routers, fields[1])
		}
	}

	for _, router := range routers {
		for _, router2 := range routers {
			if router != router2 {
				p.graph.addLink(p.area, router, router2, distance)
			}
		}
	}

	p.leftRouter()
	return nil
}
