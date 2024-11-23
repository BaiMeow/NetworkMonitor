package analysis

type BetweennessResult struct {
	Node        *Node
	Betweenness float64
}

func (g *Graph) Betweenness() []BetweennessResult {
	var passCount = make([]float64, len(g.Nodes))
	for _, node := range g.Nodes {
		paths := g.SingleSourceShortestPaths(node)
		var dstMap = make(map[int][]*Path)
		for _, path := range paths {
			dstMap[path.Dst.Id] = append(dstMap[path.Dst.Id], path)
		}
		for _, paths := range dstMap {
			for _, path := range paths {
				if len(path.Nodes) < 3 {
					continue
				}
				for _, n := range path.Nodes[1 : len(path.Nodes)-1] {
					passCount[n.Id] += 1 / float64(len(paths))
				}
			}
		}
	}

	var result []BetweennessResult
	nodeCount := float64(len(g.Nodes))
	for _, node := range g.Nodes {
		result = append(result, BetweennessResult{Node: node, Betweenness: passCount[node.Id] * 2 / (nodeCount - 1) / (nodeCount - 2)})
	}

	return result
}
