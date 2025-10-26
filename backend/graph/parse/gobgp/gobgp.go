package gobgp

import (
	"context"
	"fmt"
	"log"
	"net/netip"
	"reflect"
	"slices"

	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v3/api"
	"google.golang.org/protobuf/types/known/anypb"
)

func init() {
	parse.Register("gobgp", func(m map[string]any) (parse.Parser[*entity.BGP], error) {
		return &GoBGP{}, nil
	})
}

type GoBGP struct {
	parse.Base[*entity.BGP]
}

func (b *GoBGP) Parse(ctx context.Context, input any) (*entity.BGP, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"parse/bgp/GoBGP.Parse",
	)
	defer span.End()

	paths, ok := input.([]*apipb.Path)
	if !ok {
		log.Fatalf("invalid data type for GoBGP parser: %s", reflect.TypeOf(input).Elem().Name())
	}
	var bgp entity.BGP
	for _, p := range paths {
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
		var prefix apipb.IPAddressPrefix
		if err := p.GetNlri().UnmarshalTo(&prefix); err != nil {
			log.Println("unmarshal prefix failed:", err)
			continue
		}
		for _, se := range asPathAttr.Segments {
			numbers := se.GetNumbers()
			for i := 0; i < len(numbers)-1; i++ {
				bgp.AddAsLink(numbers[i], numbers[i+1])
			}
			bgp.AddPrefix(numbers[len(numbers)-1], netip.MustParsePrefix(fmt.Sprintf("%s/%d", prefix.GetPrefix(), prefix.GetPrefixLen())))
		}
	}

	return &bgp, nil
}
