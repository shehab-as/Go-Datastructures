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

func New() (SinglyList) {
	return SinglyList{
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (s *SinglyList) AddNodeAtHead(key int) {
	newNode := &Node{
		next: nil,
		val: key,
	}
	s.length++
	if s.head == nil && s.tail == nil {
		s.head = newNode 
		s.tail = s.head
		return
	}
	curr := s.head
	s.head = newNode
	s.head.next = curr 
}

func (s *SinglyList) AddNodeAtTail(key int) {
	newNode := &Node{ 
		next: nil,
		val: key,
	}
	s.length++
	if s.head == nil && s.tail == nil {
		s.head = newNode
		s.tail = s.head
		return
	}
	s.tail.next = newNode
	s.tail = s.tail.next
}

func (s *SinglyList) AddNode(key int, pos int) {
	if pos <= 0 { 
		return
	} else if pos == 1 {
		s.AddNodeAtHead(key)
		return
	} else if pos >= s.length {
		s.AddNodeAtTail(key)
		return
	}
	// 1->2->3
	newNode := &Node{
		next: nil,
		val: key,
	}
	s.length++
	var prev, curr *Node
	curr = s.head
	for pos != 1 && curr != nil {
		prev = curr 
		curr = curr.next
		pos--
	}
	prev.next = newNode
	newNode.next = curr
}

func (s *SinglyList) Reverse() {
	if s.head == nil && s.tail == nil {
		return
	}
	var prev, curr, fwd *Node
	prev = nil 
	curr = s.head
	for curr != nil {
		fwd = curr.next
		curr.next = prev 
		prev = curr
		curr = fwd
	}
	s.tail = s.head
	s.head = prev
}

func (s SinglyList) Display() {
	fmt.Println("List Content...")
	curr := s.head 
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println()
}