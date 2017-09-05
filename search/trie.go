package search

type trieNode struct {
	next  map[uint8]*trieNode
	match bool
	size  int
}

func NewTrieNode() *trieNode { return &trieNode{next: map[uint8]*trieNode{}} }

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie { return &Trie{} }

func (t *Trie) Contains(key string) bool {
	x := t.get(t.root, key, 0)
	if x == nil {
		return false
	}
	return true
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

func (t *Trie) Put(key string) {
	t.root = t.put(t.root, key, 0)
}

func (t *Trie) put(node *trieNode, key string, index int) *trieNode {
	if node == nil {
		node = NewTrieNode()
	}
	node.size++
	if index == len(key) {
		node.match = true
		return node
	}
	c := key[index]
	node.next[c] = t.put(node.next[c], key, index+1)
	return node
}

func (t *Trie) CountKeysWithPrefix(prefix string) int {
	node := t.get(t.root, prefix, 0)
	if node == nil {
		return 0
	}
	return node.size
}

func (t *Trie) Keys() Queue {
	return t.KeysWithPrefix("")
}

func (t *Trie) KeysWithPrefix(prefix string) Queue {
	results := NewArrayQueue()
	node := t.get(t.root, prefix, 0)
	t.collect(node, prefix, results)
	return results
}

func (t *Trie) collect(node *trieNode, prefix string, results Queue) {
	if node == nil {
		return
	}
	if node.match {
		results.Enqueue(prefix)
	}
	for k, v := range node.next {
		t.collect(v, prefix+string(k), results)
	}
}
