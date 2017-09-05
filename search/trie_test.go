package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	input := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}
	trie := NewTrie()
	for _, word := range input {
		trie.Put(word)
	}
	for _, word := range input {
		assert.True(t, trie.Contains(word))
	}
	assert.Len(t, trie.KeysWithPrefix("sh").ToArray(), 3)
	assert.Equal(t, trie.CountKeysWithPrefix("sh"), 3)
	assert.Len(t, trie.Keys().ToArray(), len(input)-1)
}
