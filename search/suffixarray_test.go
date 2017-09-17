package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuffixArray(t *testing.T) {
	sa := NewSuffixArray("banana")
	assert.EqualValues(t, sa, []*Suffix{
		{"a", 5},
		{"ana", 3},
		{"anana", 1},
		{"banana", 0},
		{"na", 4},
		{"nana", 2},
	})
}
