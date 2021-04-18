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
	mylist.Display()
	fmt.Println("Hello world!")
}