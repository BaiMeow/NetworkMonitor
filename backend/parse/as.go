package parse

import (
	"net/netip"
)

type AS struct {
	ASN     uint32         `json:"asn,omitempty"`
	Network []netip.Prefix `json:"network,omitempty"`
}

func newAS(asn uint32) *AS {
	return &AS{
		ASN: asn,
	}
}
