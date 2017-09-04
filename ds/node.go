package ds

type Node struct {
	val  int
	next *Node
}

func NewNode(val int, next *Node) *Node {
	return &Node{val, next}
}
