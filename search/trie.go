package search

import (
	"github.com/bquenin/algorithms/ds"
)

type trieNode struct {
	next  map[byte]*trieNode
	value interface{}
	size  int
}

func NewTrieNode() *trieNode { return &trieNode{next: map[byte]*trieNode{}} }

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie { return &Trie{root: NewTrieNode()} }

func (t *Trie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *Trie) get(key string) *trieNode {
	node, exists := t.root, false
	for i := 0; i < len(key); i++ {
		if node, exists = node.next[key[i]]; !exists {
			return nil
		}
	}
	return node
}

func (t *Trie) Get(key string) interface{} {
	node := t.get(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *Trie) Put(key string, value interface{}) {
	node, exists := t.root, false
	for i := 0; i < len(key); i++ {
		if _, exists = node.next[key[i]]; !exists {
			node.next[key[i]] = NewTrieNode()
		}
		node = node.next[key[i]]
		node.size++
	}
	node.value = value
}

func (t *Trie) CountKeysWithPrefix(prefix string) int {
	node := t.get(prefix)
	if node == nil {
		return 0
	}
	return node.size
}

func (t *Trie) Keys() ds.Queue {
	return t.KeysWithPrefix("")
}

func (t *Trie) KeysWithPrefix(prefix string) ds.Queue {
	results := ds.NewArrayQueue()
	node := t.get(prefix)
	t.collect(node, prefix, results)
	return results
}

func (t *Trie) collect(node *trieNode, prefix string, results ds.Queue) {
	if node == nil {
		return
	}
	if node.value != nil {
		results.Enqueue(prefix)
	}
	for k, v := range node.next {
		t.collect(v, prefix+string(k), results)
	}
}

func (t *Trie) KeysContainingAlphabet(alphabet []byte) ds.Queue {
	results := ds.NewArrayQueue()
	hash := make([]bool, 26)
	for _, letter := range alphabet {
		hash[letter-'a'] = true
	}
	t.keysContainingAlphabet(t.root, "", hash, results)
	return results
}

func (t *Trie) keysContainingAlphabet(node *trieNode, prefix string, hash []bool, results ds.Queue) {
	if node == nil {
		return
	}
	if node.value != nil {
		results.Enqueue(prefix)
	}
	for k := range node.next {
		if hash[k-'a'] {
			t.keysContainingAlphabet(node.next[k], prefix+string(k), hash, results)
		}
	}
}
