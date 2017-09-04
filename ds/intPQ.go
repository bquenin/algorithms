package ds

type IntPQ struct {
	values []int
	less   func(int, int) bool
}

func (h *IntPQ) Len() int           { return len(h.values) }
func (h *IntPQ) Swap(i, j int)      { h.values[i], h.values[j] = h.values[j], h.values[i] }
func (h *IntPQ) Less(i, j int) bool { return h.less(h.values[i], h.values[j]) }
func (h *IntPQ) Push(x interface{}) { h.values = append(h.values, x.(int)) }
func (h *IntPQ) Peek() int          { return h.values[0] }
func (h *IntPQ) Empty() bool        { return len(h.values) == 0 }
func (h *IntPQ) Pop() interface{} {
	old := h.values
	n := len(old)
	x := old[n-1]
	h.values = old[0 : n-1]
	return x
}

func NewIntMinPQ() *IntPQ { return &IntPQ{values: []int{}, less: func(i, j int) bool { return i < j }} }
func NewIntMaxPQ() *IntPQ { return &IntPQ{values: []int{}, less: func(i, j int) bool { return i > j }} }
