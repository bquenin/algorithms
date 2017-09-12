package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecursiveTrie(t *testing.T) {
	input := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}
	trie := NewRecursiveTrie()
	for i, word := range input {
		trie.Put(word, i)
	}
	for _, word := range input {
		assert.True(t, trie.Contains(word))
	}
	assert.Len(t, trie.KeysWithPrefix("sh").ToArray(), 3)
	assert.Equal(t, trie.CountKeysWithPrefix("sh"), 3)
	assert.Len(t, trie.Keys().ToArray(), len(input)-1)
}

func TestIterativeTrie(t *testing.T) {
	input := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}
	trie := NewIterativeTrie()
	for i, word := range input {
		trie.Put(word, i)
	}
	for _, word := range input {
		assert.True(t, trie.Contains(word))
	}
	assert.Len(t, trie.KeysWithPrefix("sh").ToArray(), 3)
	assert.Equal(t, trie.CountKeysWithPrefix("sh"), 3)
	assert.Len(t, trie.Keys().ToArray(), len(input)-1)
}
