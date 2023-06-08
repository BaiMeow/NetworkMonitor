package parse

import (
	"github.com/BaiMeow/OSPF-monitor/utils"
	"net/netip"
)

type BGP struct {
	AS   []*AS    `json:"as,omitempty"`
	Link []ASLink `json:"link,omitempty"`
}

func (b *BGP) AddPrefix(ASN uint32, prefix netip.Prefix) {
	target, found := utils.FindFunc(b.AS, func(as *AS) bool {
		return as.ASN == ASN
	})

	if !found {
		as := newAS(ASN)
		as.Network = append(as.Network, prefix)
		b.AS = append(b.AS, as)
		return
	}

	if utils.Find(target.Network, prefix) {
		return
	}

	target.Network = append(target.Network, prefix)
}

func (b *BGP) AddAsLink(ASN1, ASN2 uint32) {
	if ASN2 < ASN1 {
		ASN1, ASN2 = ASN2, ASN1
	}

	if utils.Find(b.Link, ASLink{Src: ASN1, Dst: ASN2}) {
		return
	}

	b.Link = append(b.Link, ASLink{Src: ASN1, Dst: ASN2})
}

func (b *BGP) Merge(bgp *BGP) {
	for _, as := range bgp.AS {
		for _, prefix := range as.Network {
			b.AddPrefix(as.ASN, prefix)
		}
	}
	for _, link := range bgp.Link {
		b.AddAsLink(link.Src, link.Dst)
	}
}
