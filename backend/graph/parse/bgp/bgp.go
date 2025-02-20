package bgp

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	apipb "github.com/osrg/gobgp/v3/api"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"net/netip"
	"reflect"
	"slices"
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

func (b *BGP) Parse(input any) (*entity.BGP, error) {
	destinations, ok := input.([]*apipb.Destination)
	if !ok {
		log.Fatalf("invalid data type for BGP parser: %s", reflect.TypeOf(input).Elem().Name())
	}
	var bgp entity.BGP

	for _, des := range destinations {
		for _, p := range des.Paths {
			idx := slices.IndexFunc(p.Pattrs, func(a *anypb.Any) bool {
				return a.GetTypeUrl() == "type.googleapis.com/apipb.AsPathAttribute"
			})
			if idx == -1 {
				continue
			}
			asPathAttrPb := p.Pattrs[idx]
			var asPathAttr apipb.AsPathAttribute
			if err := asPathAttrPb.UnmarshalTo(&asPathAttr); err != nil {
				log.Println("unmarshal ASPathAttr failed:", err)
				continue
			}
			for _, se := range asPathAttr.Segments {
				numbers := se.GetNumbers()
				if b.leftShiftCount != 0 {
					if len(numbers) <= b.leftShiftCount {
						continue
					}
					numbers = numbers[b.leftShiftCount:]
				}
				for i := 0; i < len(numbers)-1; i++ {
					bgp.AddAsLink(numbers[i], numbers[i+1])
				}
				bgp.AddPrefix(numbers[len(numbers)-1], netip.MustParsePrefix(des.GetPrefix()))
			}
		}
	}

	return &bgp, nil
}
