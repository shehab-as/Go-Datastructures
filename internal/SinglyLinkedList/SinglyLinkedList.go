package SinglyLinkedList

import (
	"fmt"
)

type Node struct {
	next *Node 
	val interface{}
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

func (s *SinglyList) AddNodeAtHead(key interface{}) {
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

func (s *SinglyList) AddNodeAtTail(key interface{}) {
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

func (s *SinglyList) AddNode(key interface{}, pos int) {
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

func (s *SinglyList) DeleteNode(pos int) {
	if pos <= 0 || pos > s.length {
		return
	}
	if pos == 1 {
		s.head = s.head.next
		return
	}
	var prev, curr *Node
	curr = s.head 
	for curr != nil && pos != 1 {
		prev = curr 
		curr = curr.next
		pos--
	}
	prev.next = curr.next
	curr.next = nil
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
		fmt.Printf("%v -> ", curr.val)
		curr = curr.next
	}
	fmt.Println()
}