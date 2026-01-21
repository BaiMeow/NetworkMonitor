package analysis

type PathBetweennessResult struct {
	Src         *Node   `json:"src"`
	Dst         *Node   `json:"dst"`
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
					passCount[[2]int{path.Nodes[i].Id, path.Nodes[i+1].Id}] += 1 / float64(len(paths))
				}
			}
		}
	}
	var results []PathBetweennessResult
	for path, count := range passCount {
		results = append(results, PathBetweennessResult{
			Src:         g.Nodes[path[0]],
			Dst:         g.Nodes[path[1]],
			Betweenness: count / float64(len(g.Nodes)*len(g.Nodes)/2),
		})
	}
	return results
}
