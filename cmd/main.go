package main

import (
	"fmt"
	"go-datastructures/internal/BinarySearchTrees"
	"go-datastructures/internal/Graphs"
	"go-datastructures/internal/SinglyLinkedList"
	"go-datastructures/internal/Stacks"
)

func main() {
	// Run_SinglyList()
	// Run_Stack()
	// Run_BST()
	Run_Graph()
}

func Run_SinglyList() {
	fmt.Println("~~~ Singly Linked List ...")
	mylist := SinglyLinkedList.New()
	mylist.AddNodeAtTail(10)
	mylist.AddNodeAtTail(20)
	mylist.AddNodeAtTail(30)
	mylist.AddNodeAtHead(100)
	mylist.AddNodeAtHead(200)
	mylist.AddNodeAtTail(40)
	mylist.AddNode(-50, 3)
	mylist.Display()
	mylist.Reverse()
	mylist.Display()
	mylist.DeleteNode(5)
	mylist.Display()
}

func Run_BST() {
	fmt.Println("~~~ Binary Search Trees ...")
	myBST := BinarySearchTrees.New()
	myBST.Add(10)
	myBST.Add(50)
	myBST.Add(30)
	myBST.Add(40)
	myBST.Add(20)
	myBST.InOrder()
}

func Run_Stack() {
	fmt.Println("~~~ Stack ...")
	myStack := Stacks.New(3)
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Push("A")
	myStack.Push("B")
	myStack.Push("C")
	myStack.Push("D")
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %v\n", myStack.Top())
}

func Run_Graph() {
	fmt.Println("~~~ Directed Graph ...")
	myGraph := Graphs.New(5)
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 3)
	myGraph.AddEdge(2, 4)
	myGraph.BFS(1)
}