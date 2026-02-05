package gobgp

import (
	"context"
	"fmt"
	"log"
	"net/netip"
	"reflect"

	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/NetworkMonitor/trace"
	apipb "github.com/osrg/gobgp/v4/api"
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
		var asPathAttr *apipb.AsPathAttribute
		for _, attr := range p.Pattrs {
			if attr.GetAsPath() != nil {
				asPathAttr = attr.GetAsPath()
			}
		}
		if asPathAttr == nil {
			log.Println("as path attribute not found")
			continue
		}
		prefix := p.GetNlri().GetPrefix()
		if prefix == nil {
			log.Println("prefix not found")
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
