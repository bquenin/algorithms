package search

import "github.com/bquenin/algorithms/ds"

type trieNode struct {
	next  map[uint8]*trieNode
	value interface{}
	size  int
}

func NewTrieNode() *trieNode { return &trieNode{next: map[uint8]*trieNode{}} }

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie { return &Trie{} }

func (t *Trie) Contains(key string) bool {
	return t.Get(key) != nil
}

func (t *Trie) Get(key string) interface{} {
	x := t.get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.value
}

func (t *Trie) get(node *trieNode, key string, index int) *trieNode {
	if node == nil {
		return nil
	}
	if index == len(key) {
		return node
	}
	c := key[index]
	return t.get(node.next[c], key, index+1)
}

func (t *Trie) Put(key string, value interface{}) {
	t.root = t.put(t.root, key, value, 0)
}

func (t *Trie) put(node *trieNode, key string, value interface{}, index int) *trieNode {
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

func (t *Trie) CountKeysWithPrefix(prefix string) int {
	node := t.get(t.root, prefix, 0)
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
	node := t.get(t.root, prefix, 0)
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
