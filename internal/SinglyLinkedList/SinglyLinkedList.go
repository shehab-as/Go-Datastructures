package SinglyLinkedList

import (
	"fmt"
)

type Node struct {
	next *Node 
	val int
}


type SinglyList struct {
	head *Node
	tail *Node
	length int
}

func New() (*SinglyList) {
	return &SinglyList{
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (s *SinglyList) AddNodeAtTail(key int) {
	newNode := &Node{ 
		next: nil,
		val: key,
	}
	if s.head == nil && s.tail == nil {
		s.head = newNode
		s.tail = s.head
		return
	}
	s.tail.next = newNode
	s.tail = s.tail.next
}

func (s *SinglyList) Display() {
	curr := s.head 
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
}