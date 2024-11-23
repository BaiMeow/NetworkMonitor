package analysis

type ClosenessResult struct {
	Node      *Node
	Closeness float64
}

func (g *Graph) Closeness() []ClosenessResult {
	var result []ClosenessResult
	for _, node := range g.Nodes {
		var sum float64
		paths := g.SingleSourceShortestPaths(node)
		for _, path := range paths {
			if path.Length == 0 {
				continue
			}
			sum += float64(path.Length)
		}
		result = append(result, ClosenessResult{Node: node, Closeness: float64(len(g.Nodes)-1) / sum})
	}
	return result
}
