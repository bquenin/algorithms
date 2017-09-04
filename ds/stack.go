package ds

type Stack struct {
	top *Node
	N   int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val int) {
	s.top = NewNode(val, s.top)
	s.N++
}

func (s *Stack) Pop() int {
	val := s.top.val
	s.top = s.top.next
	s.N--
	return val
}

func (s *Stack) Size() int {
	return s.N
}

func (s *Stack) Empty() bool {
	return s.N == 0
}
