package main

import (
	"fmt"
	"go-datastructures/internal/SinglyLinkedList"
)

func main() {

	mylist := SinglyLinkedList.New()
	mylist.AddNodeAtTail(10)
	mylist.AddNodeAtTail(20)
	mylist.AddNodeAtTail(30)
	mylist.AddNodeAtHead(100)
	mylist.AddNodeAtHead(200)
	mylist.AddNodeAtTail(40)
	mylist.AddNode(-50, 3)
	mylist.Display()
	fmt.Println()
	mylist.Reverse()
	mylist.Display()
	fmt.Println()
	fmt.Println("\nHello world!")
}