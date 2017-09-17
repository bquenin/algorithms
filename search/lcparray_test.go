package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLCPArray(t *testing.T) {
	assert.EqualValues(t, NewLCPArray("banana"), []int{1, 3, 0, 0, 2, 0})
}
