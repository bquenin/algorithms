package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRS(t *testing.T) {
	assert.Equal(t, "abra", LRS("abracadabra"))
	assert.Equal(t, "to be", LRS("to be or not to be"))
}
