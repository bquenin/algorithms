package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunningMedian(t *testing.T) {
	input := []int{12, 4, 5, 3, 8, 7}
	medians := []float64{12.0, 8.0, 5.0, 4.5, 5.0, 6.0}
	rm := NewRunningMedian()
	for i := range input {
		rm.Add(input[i])
		assert.Equal(t, medians[i], rm.Current())
	}
}
