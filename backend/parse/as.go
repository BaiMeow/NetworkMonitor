package parse

import (
	"github.com/BaiMeow/OSPF-monitor/conf"
	"net/netip"
	"strconv"
)

type AS struct {
	ASN      uint32         `json:"asn,omitempty"`
	Network  []netip.Prefix `json:"network,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

func newAS(asn uint32) *AS {
	return &AS{
		ASN:      asn,
		Metadata: conf.Metas[strconv.FormatUint(uint64(asn), 10)],
	}
}
