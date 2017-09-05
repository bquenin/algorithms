package ds

type Queue interface {
	Enqueue(val interface{})
	Dequeue() interface{}
	Size() int
	Empty() bool
	ToArray() []interface{}
}

type ArrayQueue []interface{}

func NewArrayQueue() Queue                    { return &ArrayQueue{} }
func (q *ArrayQueue) Size() int               { return len(*q) }
func (q *ArrayQueue) Empty() bool             { return len(*q) == 0 }
func (q *ArrayQueue) ToArray() []interface{}  { return *q }
func (q *ArrayQueue) Enqueue(val interface{}) { *q = append(*q, val) }
func (q *ArrayQueue) Dequeue() interface{} {
	var val interface{}
	val, *q = (*q)[0], (*q)[1:]
	return val
}

type ListQueue struct {
	first, last *Node
	N           int
}

func NewListQueue() Queue        { return &ListQueue{} }
func (q *ListQueue) Size() int   { return q.N }
func (q *ListQueue) Empty() bool { return q.N == 0 }
func (q *ListQueue) ToArray() []interface{} {
	result := make([]interface{}, q.N)
	for i := q.first; i != nil; i = i.next {
		result = append(result, i.val)
	}
	return result
}
func (q *ListQueue) Enqueue(val interface{}) {
	oldLast := q.last
	q.last = NewNode(val, nil)
	if q.Empty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.N++
}
func (q *ListQueue) Dequeue() interface{} {
	val := q.first.val
	q.first = q.first.next
	if q.Empty() {
		q.last = nil
	}
	q.N--
	return val
}
