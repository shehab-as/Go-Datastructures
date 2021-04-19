package main

import (
	"fmt"
	"go-datastructures/internal/BinarySearchTrees"
	"go-datastructures/internal/SinglyLinkedList"
	"go-datastructures/internal/Stacks"
)

func main() {
	// Run_SinglyList()
	// Run_Stack()
	Run_BST()
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
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
	myStack.Push(5)
	myStack.Push(10)
	myStack.Push(15)
	myStack.Push(20)
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
	myStack.Pop()
	fmt.Printf("Top of Stack -> %d\n", myStack.Top())
}