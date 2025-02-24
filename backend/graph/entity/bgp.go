package entity

import (
	"github.com/BaiMeow/NetworkMonitor/utils"
	"maps"
	"net/netip"
	"slices"
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
	if ASN1 == ASN2 {
		return
	}

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

func (b *BGP) Equal(b2 *BGP) bool {
	if len(b2.AS) != len(b.AS) {
		return false
	}
	if len(b2.Link) != len(b.Link) {
		return false
	}

	links := maps.Collect(func(yield func(ASLink, bool) bool) {
		slices.All(b.Link)(func(i int, link ASLink) bool {
			return yield(link, true)
		})
	})
	if slices.IndexFunc(b2.Link, func(link ASLink) bool {
		return !links[link]
	}) != -1 {
		return false
	}

	ass := maps.Collect(func(yield func(uint32, []netip.Prefix) bool) {
		slices.All(b.AS)(func(i int, as *AS) bool {
			return yield(as.ASN, as.Network)
		})
	})

	// not as not same
	return slices.IndexFunc(b2.AS, func(as *AS) bool {
		as2, ok := ass[as.ASN]
		if !ok {
			return true
		}
		if len(as2) != len(as.Network) {
			return true
		}
		// one network not equal
		return slices.IndexFunc(as.Network, func(prefix netip.Prefix) bool {
			return slices.IndexFunc(as2, func(prefix2 netip.Prefix) bool {
				return prefix == prefix2
			}) == -1
		}) != -1
	}) == -1
}
