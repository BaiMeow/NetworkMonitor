package analysis

type PathBetweennessResult struct {
	Src         *Node   `json:"src"`
	Dst         *Node   `json:"dst"`
	Cost        int     `json:"cost"`
	Betweenness float64 `json:"betweenness"`
}

func (g *Graph) PathBetweenness() []PathBetweennessResult {
	if g.ShortestPath == nil {
		g.AllSourceShortestPaths()
	}
	passCount := make(map[[2]int]float64)
	for _, dstPaths := range g.ShortestPath {
		for _, paths := range dstPaths {
			for _, path := range paths {
				for i := 0; i < len(path.Nodes)-1; i++ {
					passCount[[2]int{
						path.Nodes[i].Id,
						path.Nodes[i+1].Id,
					}] += 1 / float64(len(paths))
				}
			}
		}
	}
	var results []PathBetweennessResult
	for pair, count := range passCount {
		minCost := -1
		var samePathCount int
		for _, outEdge := range g.Nodes[pair[0]].Out {
			if outEdge.Dst.Id != pair[1] {
				continue
			}
			if minCost == -1 || minCost > outEdge.Cost {
				minCost = outEdge.Cost
				samePathCount = 1
			} else if minCost == outEdge.Cost {
				samePathCount++
			}
		}
		for range samePathCount {
			results = append(results, PathBetweennessResult{
				Src:         g.Nodes[pair[0]],
				Dst:         g.Nodes[pair[1]],
				Cost:        minCost,
				Betweenness: count / float64(len(g.Nodes)*len(g.Nodes)/2) / float64(samePathCount),
			})
		}
		for _, outEdge := range g.Nodes[pair[0]].Out {
			if outEdge.Dst.Id != pair[1] {
				continue
			}
			if outEdge.Cost > minCost {
				results = append(results, PathBetweennessResult{
					Src:         g.Nodes[pair[0]],
					Dst:         g.Nodes[pair[1]],
					Cost:        outEdge.Cost,
					Betweenness: 0,
				})
			}
		}
	}
	return results
}
