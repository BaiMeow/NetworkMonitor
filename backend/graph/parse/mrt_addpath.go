package parse

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	mrt "github.com/BaiMeow/go-mrt"
	"io"
	"net/netip"
)

func init() {
	Register("mrt-addpath", func(m map[string]any) (Parser, error) {
		return &MrtAddPath{}, nil
	})
}

var _ Parser = (*MrtAddPath)(nil)

type MrtAddPath struct {
	reader io.Reader
}

func (p *MrtAddPath) Init(input []byte) {
	p.reader = bytes.NewReader(input)
}

func (p *MrtAddPath) ParseAndMerge(drawing *Drawing) (err error) {
	var bgp BGP
	defer func() {
		if err != nil {
			return
		}
		drawing.Lock()
		drawing.BGP.Merge(&bgp)
		drawing.Unlock()
	}()

	rd := mrt.NewReader(p.reader)
	for {
		rec, err2 := rd.Next()
		if err2 != nil {
			if errors.Is(err2, io.EOF) {
				return nil
			}
			return err2
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
					return fmt.Errorf("no ASN as dst")
				}

				if !prefixAdded {
					asn := binary.BigEndian.Uint32(paths[0].Value[len(paths[0].Value)-1])
					prefix, err := netip.ParsePrefix(tb.Prefix.String())
					if err != nil {
						return err
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
