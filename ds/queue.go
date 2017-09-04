package ds

type Queue struct {
	first, last *Node
	N           int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val int) {
	oldLast := q.last
	q.last = NewNode(val, nil)
	if q.Empty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
	q.N++
}

func (q *Queue) Dequeue() int {
	val := q.first.val
	q.first = q.first.next
	if q.Empty() {
		q.last = nil
	}
	q.N--
	return val
}

func (q *Queue) Size() int {
	return q.N
}

func (q *Queue) Empty() bool {
	return q.N == 0
}
