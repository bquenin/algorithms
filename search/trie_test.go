package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	input := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}
	trie := NewTrie()
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

func TestDictionaryAndLetters(t *testing.T) {
	dictionary := []string{"go", "bat", "me", "eat", "goal", "boy", "run"}
	letters := []byte{'e', 'o', 'b', 'a', 'm', 'g', 'l'}
	trie := NewTrie()
	for i, word := range dictionary {
		trie.Put(word, i)
	}
	matched := trie.KeysContainingLetters(letters).ToArray()
	assert.Len(t, matched, 3)
	assert.Contains(t, matched, "me")
	assert.Contains(t, matched, "go")
	assert.Contains(t, matched, "goal")
}
