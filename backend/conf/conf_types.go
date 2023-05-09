package conf

type Probe struct {
	Parse Parser
	Fetch Fetcher
}

type Parser map[string]any

func (p Parser) Type() string {
	return p["type"].(string)
}

type Fetcher map[string]any

func (f Fetcher) Type() string {
	return f["type"].(string)
}
