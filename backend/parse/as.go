package parse

import (
	"net/netip"
)

type AS struct {
	ASN     uint32         `json:"asn"`
	Network []netip.Prefix `json:"network"`
}

func newAS(asn uint32) *AS {
	return &AS{
		ASN: asn,
	}
}
