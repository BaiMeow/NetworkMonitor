package parse

type Link struct {
	Src  string `json:"src,omitempty"`
	Dst  string `json:"dst,omitempty"`
	Cost int    `json:"cost,omitempty"`
}

func newLink(src, dst string, cost int) Link {
	return Link{src, dst, cost}
}
