package BinarySearchTrees

import "fmt"


type Node struct {
	left *Node 
	right *Node 
	val int
}

type BST struct {
	root *Node 
}

func New() (BST) {
	return BST{
		root: nil,
	}
}

func (t BST) _addNode(key int, root *Node) (*Node) {
	if root == nil {
		return &Node{
			left: nil,
			right: nil,
			val: key,
		}
	}
	if key < root.val {
		root.left = t._addNode(key, root.left)
	} else {
		root.right = t._addNode(key, root.right)
	}
	return root
}

func (t *BST) Add(key int) {
	t.root = t._addNode(key, t.root)
}

func (t *BST) _inOrder(root *Node) {
	if root == nil {
		return
	}
	t._inOrder(root.left)
	fmt.Printf("%d ", root.val)
	t._inOrder(root.right)
}
func (t BST) InOrder() {
	fmt.Println("InOrder Sequence...")
	t._inOrder(t.root)
}