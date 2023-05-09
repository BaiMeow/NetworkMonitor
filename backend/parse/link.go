package parse

import "strings"

type Link struct {
	Src  string `json:"src,omitempty"`
	Dst  string `json:"dst,omitempty"`
	Cost int    `json:"cost,omitempty"`
}

func newLink(src, dst string, cost int) Link {
	if strings.Compare(src, dst) == 1 {
		src, dst = dst, src
	}
	return Link{src, dst, cost}
}
