package analysis

type ClosenessResult struct {
	Node      *Node
	Closeness float64
}

func (g *Graph) Closeness() []ClosenessResult {
	var result []ClosenessResult
	for _, node := range g.Nodes {
		paths := g.SingleSourceShortestPaths(node)
		distance := make([]int, len(g.Nodes))
		for _, path := range paths {
			distance[path.Dst.Id] = path.Length
		}
		var sum int
		for i, v := range distance {
			if v == 0 && i != node.Id {
				sum = -1
				break
			}
			sum += v
		}
		if sum == -1 {
			// graph not full connected, not supported
			result = append(result, ClosenessResult{Node: node, Closeness: -1})
		}
		result = append(result, ClosenessResult{Node: node, Closeness: float64(len(g.Nodes)-1) / float64(sum)})
	}
	return result
}
