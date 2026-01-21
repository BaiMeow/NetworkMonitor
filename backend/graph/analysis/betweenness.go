package analysis

import (
	"slices"
)

type BetweennessResult struct {
	Node        *Node
	Betweenness float64
}

func (g *Graph) Betweenness() []BetweennessResult {
	if g.ShortestPath == nil {
		g.AllSourceShortestPaths()
	}
	var passCount = make([]float64, len(g.Nodes))
	for _, node := range g.Nodes {
		for _, paths := range g.ShortestPath[node.Id] {
			paths = slices.DeleteFunc(slices.Clone(paths), func(p *Path) bool {
				return len(p.Nodes) < 3
			})
			if len(paths) == 0 {
				continue
			}
			for _, path := range paths {
				for _, n := range path.Nodes[1 : len(path.Nodes)-1] {
					passCount[n.Id] += 1 / float64(len(paths))
				}
			}
		}
	}

	var result []BetweennessResult
	nodeCount := float64(len(g.Nodes))
	for _, node := range g.Nodes {
		result = append(result, BetweennessResult{Node: node, Betweenness: passCount[node.Id] / (nodeCount - 1) / (nodeCount - 2)})
	}

	return result
}
