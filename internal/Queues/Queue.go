package Queues

import "fmt"

type Queue struct {
	front 		int 
	back 		int 
	queue 		[]interface{}
	size		int
	capacity	int
}

func New(capacity int) (Queue) {
	return Queue{
		front: 1,
		back: 0,
		size: 0,
		capacity: capacity,
		queue: make([]interface{}, capacity),
	}
}

func (q *Queue) Enqueue(key interface{}) {
	if q.IsFull() {
		return
	}
	q.back = (q.back + 1) % q.capacity
	q.queue[q.back] = key
	fmt.Printf("Pushed %v to the queue.\n", key)
	q.size++
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		return
	}
	key := q.queue[q.front]
	fmt.Printf("Popped %v from the queue.\n", key)
	q.queue[q.front] = nil 
	q.front = (q.front + 1) % q.capacity
	q.size--
}

func (q *Queue) Front() (key interface{}) {
	if q.IsEmpty() {
		return nil
	}
	return q.queue[q.front]
}

func (q *Queue) IsEmpty() (bool) {
	return q.size == 0
}

func (q *Queue) IsFull() (bool) {
	return q.size == q.capacity
}