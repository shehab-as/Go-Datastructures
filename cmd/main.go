package main

import (
	"fmt"
	"go-datastructures/internal/BinarySearchTrees"
	"go-datastructures/internal/Graphs"
	"go-datastructures/internal/Queues"
	"go-datastructures/internal/SinglyLinkedList"
	"go-datastructures/internal/Stacks"
	"go-datastructures/internal/Tries"
)

func main() {
	// Run_SinglyList()
	// Run_Stack()
	// Run_BST()
	// Run_Graph()
	// Run_Queue()
	Run_Trie()
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

func Run_Queue() {
	myQueue := Queues.New(100)
	myQueue.Enqueue("Customer 1")
	myQueue.Enqueue("Customer 2")
	myQueue.Enqueue("Customer 3")
	myQueue.Enqueue("Customer 4")
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
	myQueue.Dequeue()
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
	myQueue.Dequeue()
	myQueue.Enqueue("Customer 5")
	fmt.Printf("Front of the Queue - %v.\n", myQueue.Front())
}

func Run_Graph() {
	fmt.Println("~~~ Directed Graph ...")
	myGraph := Graphs.New(5)
	myGraph.AddEdge(1, 2)
	myGraph.AddEdge(1, 4)
	myGraph.AddEdge(1, 5)
	myGraph.AddEdge(2, 3)
	myGraph.AddEdge(3, 4)
	myGraph.AddEdge(5, 6)
	myGraph.AddEdge(6, 7)
	myGraph.AddEdge(7, 8)
	myGraph.AddEdge(8, 9)
	myGraph.AddEdge(3, 10)

	fmt.Println("BFS...")
	myGraph.BFS(1)
	fmt.Println("DFS...")
	myGraph.DFS(1)
}
/*        2 - - > 3 -> 10
         /        |
		1  -> 4 <-
         \ 
		  5 -> 6 -> 7 -> 8 -> 9
*/

func Run_Trie() {
	myTrie := Tries.New()
	myTrie.AddKey("She")
	myTrie.AddKey("Shehab")
	fmt.Printf("\"%s\" found ? %v\n", "She", myTrie.SearchKey("She"))
	fmt.Printf("\"%s\" found ? %v\n", "Shehab", myTrie.SearchKey("Shehab"))
	fmt.Printf("\"%s\" found ? %v\n", "he", myTrie.SearchKey("he"))
}