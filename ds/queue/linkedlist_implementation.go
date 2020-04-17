package queue

import (
	"../linkedlist"

	"errors"
)

type LinkedQueue struct {
	list linkedlist.LinkedList
}

func (queue *LinkedQueue) Enqueue(value interface{}) {
	queue.list.Push(value)
}

func (queue *LinkedQueue) Dequeue() (value interface{}, err error) {
	if queue.list.Length == 0 {
		return nil, errors.New("queue is empty")
	}
	value = queue.list.Shift()
	return
}

func (queue *LinkedQueue) Length() int {
	return queue.list.Length
}
