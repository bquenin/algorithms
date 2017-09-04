package sort

import (
	"container/heap"

	"github.com/bquenin/algorithms/ds"
)

type RunningMedian struct {
	low, high *ds.IntPQ
}

func NewRunningMedian() *RunningMedian {
	return &RunningMedian{low: ds.NewIntMaxPQ(), high: ds.NewIntMinPQ()}
}

func (rm *RunningMedian) Add(val int) {
	heap.Push(rm.high, val)       // Push value into the higher end
	val = heap.Pop(rm.high).(int) // Pop the minimum of the higher end
	heap.Push(rm.low, val)        // Push this value into the lower end
	val = heap.Pop(rm.low).(int)  // Pop the maximum of the lower end
	if rm.low.Len() >= rm.high.Len() {
		heap.Push(rm.high, val)
	} else {
		heap.Push(rm.low, val)
	}
}

func (rm *RunningMedian) Current() float64 {
	if rm.high.Len() > rm.low.Len() {
		return float64(rm.high.Peek())
	} else {
		return float64(rm.high.Peek()+rm.low.Peek()) / 2
	}
}
