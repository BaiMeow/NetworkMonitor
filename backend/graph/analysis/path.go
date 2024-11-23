package analysis

type Path struct {
	Src    *Node
	Dst    *Node
	Nodes  []*Node
	Length int
}
