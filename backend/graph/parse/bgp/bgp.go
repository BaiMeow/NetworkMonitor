package bgp

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"slices"

	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"github.com/osrg/gobgp/v4/pkg/apiutil"
	"github.com/osrg/gobgp/v4/pkg/packet/bgp"
)

func init() {
	parse.Register("bgp", func(m map[string]any) (parse.Parser[*entity.BGP], error) {
		var bgp BGP
		if m["left-shift"] != nil {
			if leftShiftInt, ok := m["left-shift"].(int); ok {
				bgp.leftShiftCount = leftShiftInt
			} else {
				return nil, fmt.Errorf("left-shift set but not int")
			}
		}
		return &bgp, nil
	})
}

type BGP struct {
	parse.Base[*entity.BGP]
	leftShiftCount int
}

func (b *BGP) Parse(ctx context.Context, input any) (*entity.BGP, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"parse/bgp/BGP.Parse",
	)
	defer span.End()
	inputArr, ok := input.([][2]any)
	if !ok {
		log.Fatalf("invalid data type for BGP parser: %s", reflect.TypeOf(input).Elem().Name())
	}
	var gr entity.BGP

	for _, p := range inputArr {
		nlri, ok := p[0].(bgp.NLRI)
		if !ok {
			return nil, fmt.Errorf("invalid prefix: %s", p[0])
		}
		prefix, ok := nlri.(*bgp.IPAddrPrefix)
		if !ok {
			return nil, fmt.Errorf("invalid prefix: %s", nlri)
		}
		paths, ok := p[1].([]*apiutil.Path)
		for _, p := range paths {
			idx := slices.IndexFunc(p.Attrs, func(a bgp.PathAttributeInterface) bool {
				_, ok := a.(*bgp.PathAttributeAsPath)
				return ok
			})
			if idx == -1 {
				continue
			}
			asPathAttrPb := p.Attrs[idx].(*bgp.PathAttributeAsPath)
			for _, se := range asPathAttrPb.Value {
				numbers := se.GetAS()
				if b.leftShiftCount != 0 {
					if len(numbers) <= b.leftShiftCount {
						continue
					}
					numbers = numbers[b.leftShiftCount:]
				}
				for i := 0; i < len(numbers)-1; i++ {
					gr.AddAsLink(numbers[i], numbers[i+1])
				}
				gr.AddPrefix(numbers[len(numbers)-1], prefix.Prefix)
			}
		}
	}
	return &gr, nil
}
