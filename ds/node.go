package ds

type Node struct {
	val  interface{}
	next *Node
}

func NewNode(val interface{}, next *Node) *Node { return &Node{val, next} }
