package search

import (
	"github.com/bquenin/algorithms/ds"
)

type Trie interface {
	Contains(key string) bool
	Get(key string) interface{}
	Put(key string, value interface{})
	CountKeysWithPrefix(prefix string) int
	Keys() ds.Queue
	KeysWithPrefix(prefix string) ds.Queue
}

type trieNode struct {
	next  map[byte]*trieNode
	value interface{}
	size  int
}

func NewTrieNode() *trieNode { return &trieNode{next: map[byte]*trieNode{}} }

type IterativeTrie struct {
	root *trieNode
}

func NewIterativeTrie() *IterativeTrie { return &IterativeTrie{root: NewTrieNode()} }

func (t *IterativeTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *IterativeTrie) get(key string) *trieNode {
	node, exists := t.root, false
	for i := 0; i < len(key); i++ {
		if node, exists = node.next[key[i]]; !exists {
			return nil
		}
	}
	return node
}

func (t *IterativeTrie) Get(key string) interface{} {
	node := t.get(key)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *IterativeTrie) Put(key string, value interface{}) {
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

func (t *IterativeTrie) CountKeysWithPrefix(prefix string) int {
	node := t.get(prefix)
	if node == nil {
		return 0
	}
	return node.size
}

func (t *IterativeTrie) Keys() ds.Queue {
	return t.KeysWithPrefix("")
}

func (t *IterativeTrie) KeysWithPrefix(prefix string) ds.Queue {
	results := ds.NewArrayQueue()
	node := t.get(prefix)
	t.collect(node, prefix, results)
	return results
}

func (t *IterativeTrie) collect(node *trieNode, prefix string, results ds.Queue) {
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

type RecursiveTrie struct {
	root *trieNode
}

func NewRecursiveTrie() *RecursiveTrie { return &RecursiveTrie{} }

func (t *RecursiveTrie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *RecursiveTrie) Get(key string) interface{} {
	node := t.get(t.root, key, 0)
	if node == nil {
		return nil
	}
	return node.value
}

func (t *RecursiveTrie) get(node *trieNode, key string, index int) *trieNode {
	if node == nil {
		return nil
	}
	if index == len(key) {
		return node
	}
	c := key[index]
	return t.get(node.next[c], key, index+1)
}

func (t *RecursiveTrie) Put(key string, value interface{}) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *RecursiveTrie) put(node *trieNode, key string, value interface{}, index int) *trieNode {
	if node == nil {
		node = NewTrieNode()
	}
	node.size++
	if index == len(key) {
		node.value = value
		return node
	}
	c := key[index]
	node.next[c] = t.put(node.next[c], key, value, index+1)
	return node
}

func (t *RecursiveTrie) CountKeysWithPrefix(prefix string) int {
	node := t.get(t.root, prefix, 0)
	if node == nil {
		return 0
	}
	return node.size
}

func (t *RecursiveTrie) Keys() ds.Queue {
	return t.KeysWithPrefix("")
}

func (t *RecursiveTrie) KeysWithPrefix(prefix string) ds.Queue {
	results := ds.NewArrayQueue()
	node := t.get(t.root, prefix, 0)
	t.collect(node, prefix, results)
	return results
}

func (t *RecursiveTrie) collect(node *trieNode, prefix string, results ds.Queue) {
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
