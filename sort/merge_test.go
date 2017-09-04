package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	input := []int{12, 2, 4, 8, 8, 8, 9, 43, 1, 5, 4, 3, 9, 76, 5, 4}
	MergeSort(input)
	assert.True(t, sort.IntsAreSorted(input))
}
