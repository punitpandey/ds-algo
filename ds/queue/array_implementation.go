package queue

import "errors"

type ArrayQueue struct {
	list   []interface{}
	length int
}

func (queue *ArrayQueue) Enqueue(value interface{}) {
	queue.list = append(queue.list, value)
	queue.length++
}

func (queue *ArrayQueue) Dequeue() (value interface{}, err error) {
	if queue.length == 0 {
		return nil, errors.New("queue is empty")
	}
	value = queue.list[0]
	queue.list = queue.list[1:]
	queue.length--
	return
}

func (queue *ArrayQueue) Length() (value int) {
	return queue.length
}
