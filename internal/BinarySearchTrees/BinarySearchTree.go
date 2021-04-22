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

func (t *BST) Add(key int) {
	fmt.Printf("Adding key: %d\n", key)
	t.root = t._addNode(key, t.root)
}

func (t BST) _addNode(key int, root *Node) *Node {
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

func (t *BST) Delete(key int) {
	fmt.Printf("Deleting key: %d\n", key)
	t.root = t._deleteNode(key, t.root)
}

func (t BST) _deleteNode(key int, root *Node) *Node {
	if root == nil { return nil }
	if key < root.val {
		root.left = t._deleteNode(key, root.left)
	} else if key > root.val {
		root.right = t._deleteNode(key, root.right)
	} else {
		if root.left == nil {
			rightChild := root.right 
			root = nil 
			return rightChild
		}
		if root.right == nil {
			leftChild := root.left
			root = nil 
			return leftChild
		}
		successor := root.right 
		for successor.left != nil {
			successor = successor.left
		}
		root.val, successor.val = successor.val, root.val
		root.right = t._deleteNode(key, root.right)
	}
	return root
}

func (t BST) _searchNode(key int, root *Node) bool {
	if(root == nil) { return false }
	if(key == root.val) { return true }
	if key > root.val { 
		return t._searchNode(key, root.right) 
		} 
	return t._searchNode(key, root.left)
}

func (t BST) Search(key int) bool {
	return t._searchNode(key, t.root)
}

func (t BST) InOrder() {
	fmt.Println("InOrder Sequence...")
	t._inOrder(t.root)
	fmt.Println()
}

func (t BST) PreOrder() {
	fmt.Println("PreOrder Sequence...")
	t._preOrder(t.root)
	fmt.Println()
}

func (t BST) PostOrder() {
	fmt.Println("PostOrder Sequence...")
	t._postOrder(t.root)
	fmt.Println()
}

func (t *BST) _inOrder(root *Node) {
	if root == nil { return }
	t._inOrder(root.left)
	fmt.Printf("%d ", root.val)
	t._inOrder(root.right)
}

func (t BST) _preOrder(root *Node) {
	if root == nil { return }
	fmt.Printf("%d ", root.val)
	t._preOrder(root.left)
	t._preOrder(root.right)
}

func (t BST) _postOrder(root *Node) {
	if root == nil { return }
	t._preOrder(root.left)
	t._preOrder(root.right)
	fmt.Printf("%d ", root.val)
}