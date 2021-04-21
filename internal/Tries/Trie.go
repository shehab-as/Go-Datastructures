package Tries

import "fmt"

// Representing ASCI table 256 chars.
type Node struct {
	children	[]*Node
	isLeaf		bool
}

func NewNode() (*Node) {
	return &Node{
		children: make([]*Node, 256),
		isLeaf: false,
	}
}

type Trie struct {
	root 	*Node
}

func New() (Trie) {
	return Trie{
		root: NewNode(),
	}
}

func (t *Trie) SearchKey(key string) (bool) {
	fmt.Printf("Searching Key - %s\n", key)
	if t.root == nil {
		return false
	}

	curr := t.root
	for _, c := range key {
		k := c - '0'
		if curr.children[k] == nil {
			return false
		}
		curr = curr.children[k]
	}
	return curr != nil && curr.isLeaf
}

func (t *Trie) AddKey(key string) {
	fmt.Printf("Adding Key - %s\n", key)
	curr := t.root
	for _, c := range key {
		k := c - '0'
		if curr.children[k] == nil {
			curr.children[k] = NewNode()
		}
		curr = curr.children[k]
	}
	curr.isLeaf = true
}