package analysis

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_SingleSourceShortestPaths(t *testing.T) {
	g := &Graph{}
	for i := 0; i < 6; i++ {
		g.Nodes = append(g.Nodes, &Node{Id: i})
	}
	addLink := func(src, dst *Node) {
		src.Out = append(src.Out, &Edge{Src: src, Dst: dst, Cost: 1})
		dst.Out = append(dst.Out, &Edge{Src: dst, Dst: src, Cost: 1})
	}
	var sb strings.Builder
	for i := 0; i < 30; i++ {
		src := rand.Intn(6)
		dst := rand.Intn(5)
		if dst == src {
			dst = 5
		}
		addLink(g.Nodes[src], g.Nodes[dst])
		sb.WriteString(fmt.Sprintf("%d --> %d\n", src, dst))
	}
	fmt.Println(sb.String())
	paths := g.SingleSourceShortestPaths(g.Nodes[0])
	printPath := func(p *Path) {
		t.Logf("Path from %d to %d: %v", p.Src.Id, p.Dst.Id, p.Length)
		sb.Reset()
		for _, n := range p.Nodes {
			sb.WriteString(fmt.Sprintf("%d --> ", n.Id))
		}
		t.Log(sb.String()[:sb.Len()-5])
	}
	for _, p := range paths {
		printPath(p)
	}
}

func TestGraph_SingleSourceShortestPaths2(t *testing.T) {
	g := &Graph{}
	for i := 0; i < 2; i++ {
		g.Nodes = append(g.Nodes, &Node{Id: i})
	}
	addLink := func(src, dst *Node) {
		src.Out = append(src.Out, &Edge{Src: src, Dst: dst, Cost: 1})
		dst.Out = append(dst.Out, &Edge{Src: dst, Dst: src, Cost: 1})
	}
	addLink(g.Nodes[0], g.Nodes[1])
	addLink(g.Nodes[0], g.Nodes[1])
	paths := g.SingleSourceShortestPaths(g.Nodes[0])
	assert.Equal(t, 2, len(slices.DeleteFunc(paths, func(p *Path) bool {
		return p.Dst.Id != 1
	})), "There should be 2 paths from 0 to 1")
}

func TestGraph_All(t *testing.T) {
	g := &Graph{}
	for i := 0; i < 2; i++ {
		g.Nodes = append(g.Nodes, &Node{Id: i})
	}
	addLink := func(src, dst *Node) {
		src.Out = append(src.Out, &Edge{Src: src, Dst: dst, Cost: 1})
		dst.Out = append(dst.Out, &Edge{Src: dst, Dst: src, Cost: 1})
	}
	addLink(g.Nodes[0], g.Nodes[1])
	addLink(g.Nodes[0], g.Nodes[1])
	g.Betweenness()
	g.Closeness()
	g.PathBetweenness()
}

func BenchmarkGraph_AllSourceShortestPaths(b *testing.B) {
	const M = 300
	const N = 60
	g := &Graph{}
	for i := 0; i < N; i++ {
		g.Nodes = append(g.Nodes, &Node{Id: i})
	}
	addLink := func(src, dst *Node) {
		src.Out = append(src.Out, &Edge{Src: src, Dst: dst, Cost: 1})
		dst.Out = append(dst.Out, &Edge{Src: dst, Dst: src, Cost: 1})
	}
	for i := 0; i < M; i++ {
		src := rand.Intn(N)
		dst := rand.Intn(N - 1)
		if dst == src {
			dst = N - 1
		}
		addLink(g.Nodes[src], g.Nodes[dst])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.AllSourceShortestPaths()
	}
}
