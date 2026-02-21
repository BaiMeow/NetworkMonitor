package analysis

import (
	"slices"

	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
)

type Node struct {
	// Must Count From 0
	Id  int
	Tag map[string]any
	Out []*Edge
	In  []*Edge
}

type Edge struct {
	// Must Count From 0
	Id   int
	Src  *Node
	Dst  *Node
	Cost int
}

type Graph struct {
	// no directional graph
	bidirectional bool
	Nodes         []*Node
	ShortestPath  map[int]map[int][]*Path
}

func ConvertFromBGP(bgp *entity.BGP) *Graph {
	var (
		nodeIdReleaser IdReleaser
		edgeIdReleaser IdReleaser
	)
	g := &Graph{
		bidirectional: true,
	}
	for _, as := range bgp.AS {
		node := &Node{Id: nodeIdReleaser.Next()}
		node.Tag = map[string]any{"asn": as.ASN}
		g.Nodes = append(g.Nodes, node)
	}
	for _, link := range bgp.Link {
		var srcNode *Node
		var dstNode *Node
		for _, node := range g.Nodes {
			if node.Tag["asn"].(uint32) == link.Src {
				srcNode = node
			}
			if node.Tag["asn"].(uint32) == link.Dst {
				dstNode = node
			}
		}
		if srcNode == nil || dstNode == nil {
			continue
		}
		edge := &Edge{Id: edgeIdReleaser.Next(), Src: srcNode, Dst: dstNode, Cost: 1}
		srcNode.Out = append(srcNode.Out, edge)
		dstNode.In = append(dstNode.In, edge)
		edgeRe := &Edge{Id: edgeIdReleaser.Next(), Src: dstNode, Dst: srcNode, Cost: 1}
		dstNode.Out = append(dstNode.Out, edgeRe)
		srcNode.In = append(srcNode.In, edgeRe)
	}
	return g
}

func ConvertFromOSPF(ospf *entity.OSPF) *Graph {
	var (
		nodeIdReleaser IdReleaser
		edgeIdReleaser IdReleaser
	)
	g := &Graph{}
	for _, area := range *ospf {
		for _, router := range area.Router {
			node := &Node{Id: nodeIdReleaser.Next()}
			node.Tag = map[string]any{"routerId": router.RouterId}
			g.Nodes = append(g.Nodes, node)
		}
		for _, link := range area.Links {
			var (
				srcNode *Node
				dstNode *Node
			)
			for _, node := range g.Nodes {
				if node.Tag["routerId"].(string) == link.Src {
					srcNode = node
				}
				if node.Tag["routerId"].(string) == link.Dst {
					dstNode = node
				}
			}
			if srcNode == nil || dstNode == nil {
				continue
			}
			edge := &Edge{Id: edgeIdReleaser.Next(), Src: srcNode, Dst: dstNode, Cost: link.Cost}
			srcNode.Out = append(srcNode.Out, edge)
			dstNode.In = append(dstNode.In, edge)
		}
	}
	return g
}

func (g *Graph) FindNode(filter func(node *Node) bool) *Node {
	for _, node := range g.Nodes {
		if filter(node) {
			return node
		}
	}
	return nil
}

func (g *Graph) SingleSourceShortestPaths(src *Node) []*Path {
	findN := 1
	maxLen := 0
	var (
		paths   = make([][]*Path, len(g.Nodes))
		visited = make([]bool, len(g.Nodes))
	)

	paths[src.Id] = []*Path{{
		Src:    src,
		Dst:    src,
		Nodes:  []*Node{src},
		Length: 0,
	}}

	edgeQueue := pq.NewWith(func(a, b interface{}) int {
		// If an edge in the edgeQueue, there must be a path to the edge's src node
		return paths[a.(*Edge).Src.Id][0].Length + a.(*Edge).Cost - (paths[b.(*Edge).Src.Id][0].Length + b.(*Edge).Cost)
	})

	// Enqueue all edges from the source node
	for _, v := range g.FindNode(func(node *Node) bool {
		return src.Id == node.Id
	}).Out {
		edgeQueue.Enqueue(v)
	}

	// Mark the source node as its edges has been enqueued
	visited[src.Id] = true

	for {
		e, ok := edgeQueue.Dequeue()
		if !ok {
			break
		}
		edge := e.(*Edge)

		pathLen := paths[edge.Src.Id][0].Length + edge.Cost
		if len(paths[edge.Dst.Id]) == 0 {
			// shortest path to new node
			paths[edge.Dst.Id] = make([]*Path, len(paths[edge.Src.Id]))
			for i, frontPath := range paths[edge.Src.Id] {
				paths[edge.Dst.Id][i] = &Path{
					Src:    src,
					Dst:    edge.Dst,
					Nodes:  append(slices.Clone(frontPath.Nodes), edge.Dst),
					Length: pathLen,
				}
			}
			findN++
			if pathLen > maxLen {
				maxLen = pathLen
			}
		} else if pathLen == paths[edge.Dst.Id][0].Length {
			// already have a path with the same length
			for _, frontPath := range paths[edge.Src.Id] {
				paths[edge.Dst.Id] = append(paths[edge.Dst.Id], &Path{
					Src:    src,
					Dst:    edge.Dst,
					Nodes:  append(slices.Clone(frontPath.Nodes), edge.Dst),
					Length: pathLen,
				})
			}
		}

		if !visited[edge.Dst.Id] {
			visited[edge.Dst.Id] = true
			for _, in := range edge.Dst.Out {
				edgeQueue.Enqueue(in)
			}
		}

		if findN == len(g.Nodes) && pathLen > maxLen || edgeQueue.Empty() {
			break
		}
	}

	var allPaths []*Path
	for _, path := range paths {
		if len(path) == 0 {
			continue
		}
		allPaths = append(allPaths, path...)
	}
	return allPaths
}

func (g *Graph) AllSourceShortestPaths() {
	if g.ShortestPath == nil {
		g.ShortestPath = make(map[int]map[int][]*Path)
	}
	for _, node := range g.Nodes {
		if _, ok := g.ShortestPath[node.Id]; !ok {
			g.ShortestPath[node.Id] = make(map[int][]*Path)
		}
		paths := g.SingleSourceShortestPaths(node)
		for _, path := range paths {
			g.ShortestPath[node.Id][path.Dst.Id] = append(g.ShortestPath[node.Id][path.Dst.Id], path)
		}
	}
}

func (g *Graph) GetCloseCost(src int, dst int) int {
	if src < 0 || src >= len(g.Nodes) {
		return -1
	}
	paths := g.ShortestPath[src][dst]
	if len(paths) == 0 {
		return -1
	}
	return paths[0].Length
}
