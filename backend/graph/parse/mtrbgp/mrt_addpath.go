package mtrbgp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/BaiMeow/go-mrt"
	"io"
	"log"
	"net/netip"
	"reflect"
)

func init() {
	parse.Register("mrt-addpath", func(m map[string]any) (parse.Parser[*entity.BGP], error) {
		return &MrtAddPath{}, nil
	})
}

var _ parse.Parser[*entity.BGP] = (*MrtAddPath)(nil)

type MrtAddPath struct {
	parse.Base[*entity.BGP]
}

func (p *MrtAddPath) Parse(input any) (*entity.BGP, error) {
	data, ok := input.([]byte)
	if !ok {
		log.Fatalf("invalid input data type: %s\n", reflect.TypeOf(input).Elem())
	}
	reader := bytes.NewReader(data)

	var bgp entity.BGP

	rd := mrt.NewReader(reader)
	for {
		rec, err2 := rd.Next()
		if err2 != nil {
			if errors.Is(err2, io.EOF) {
				return &bgp, nil
			}
			return nil, err2
		}
		tb, ok := rec.(*mrt.TableDumpV2RIB)
		if !ok {
			continue
		}
		prefixAdded := false

		for _, v := range tb.RIBEntries {
			for _, attr := range v.BGPAttributes {
				paths, ok := attr.Value.(mrt.BGPPathAttributeASPath)
				if !ok {
					continue
				}
				if len(paths) == 0 {
					continue
				}

				// at least contains an ASN as dst
				if len(paths[0].Value) == 0 {
					return nil, fmt.Errorf("no ASN as dst")
				}

				if !prefixAdded {
					asn := binary.BigEndian.Uint32(paths[0].Value[len(paths[0].Value)-1])
					prefix, err := netip.ParsePrefix(tb.Prefix.String())
					if err != nil {
						return nil, err
					}
					bgp.AddPrefix(asn, prefix)
					prefixAdded = true
				}

				for _, path := range paths {
					ASPath := path.Value
					if len(ASPath) < 2 {
						continue
					}
					former := binary.BigEndian.Uint32(ASPath[0])
					for _, asn := range ASPath[1:] {
						latter := binary.BigEndian.Uint32(asn)
						bgp.AddAsLink(former, latter)
						former = latter
					}
				}
			}
		}
	}
}
