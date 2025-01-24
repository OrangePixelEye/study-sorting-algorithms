package structures

import "errors"

// FIFO

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds an element to the back of the queue.
func (q *Queue) Enqueue(data int) {
	newNode := &QueueNode{data: data}
	if q.head == nil {
		q.head = newNode
	} else {
		q.tail.next = newNode
	}
	q.tail = newNode
}

// Dequeue removes and returns the element at the front of the queue.
func (q *Queue) Dequeue() (int, error) {
	if q.head == nil {
		return 0, errors.New("Queue is empty")
	}
	item := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return item, nil
}
