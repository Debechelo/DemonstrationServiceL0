package service

import (
	"sync"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type Queue struct {
	lock  sync.Mutex
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(item interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	newNode := &Node{Value: item, Next: nil}

	if q.front == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.Next = newNode
		q.rear = newNode
	}
}

func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.front == nil {
		return nil
	}

	item := q.front.Value
	q.front = q.front.Next

	if q.front == nil {
		q.rear = nil
	}

	return item
}
