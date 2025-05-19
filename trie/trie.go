package trie

import (
	"errors"
)

// transform the val to offset in array
type TransformFunc func(val interface{}) int

// outside function only can access this struct
type Trie struct {
	root      *trieNode
	transform TransformFunc
	size      int
}

// internal struct
type trieNode struct {
	// val is option, no need in some case, such as string match
	val    interface{}
	isLeaf bool
	nexts  []*trieNode
}

func newTrieNode(size int) *trieNode {
	return &trieNode{
		isLeaf: false,
		val:    nil,
		nexts:  make([]*trieNode, size),
	}
}

func Init(size int, transFunc TransformFunc) *Trie {
	return &Trie{
		size:      size,
		root:      newTrieNode(size),
		transform: transFunc,
	}
}

func (t *Trie) Get(path []interface{}) (interface{}, error) {
	curr := t.root
	for _, item := range path {
		pos := t.transform(item)
		if pos >= t.size || curr.nexts[pos] == nil {
			return nil, errors.New("item not existed")
		}
		// curr = next
		curr = curr.nexts[pos]
	}
	if curr.isLeaf {
		return curr.val, nil
	}
	return nil, errors.New("item not existed")
}

func (t *Trie) Add(path []interface{}, val interface{}) error {
	curr := t.root
	for _, item := range path {
		pos := t.transform(item)
		if pos >= t.size {
			return errors.New("transform func got out of index")
		}

		if curr.nexts[pos] == nil {
			curr.nexts[pos] = newTrieNode(t.size)
		}
		curr = curr.nexts[pos]
	}
	curr.isLeaf = true
	curr.val = val
	return nil
}

func (t *Trie) Remove(path []interface{}) {
	curr := t.root
	for _, item := range path {
		pos := t.transform(item)
		if pos >= len(curr.nexts) || curr.nexts[pos] == nil {
			return
		}
		curr = curr.nexts[pos]
	}
	// remove the leaf mark, lazy delete
	curr.isLeaf = false
}
