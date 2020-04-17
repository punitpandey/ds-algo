package stack

import (
	"../linkedlist"

	"errors"
)

type LinkedStack struct {
	list linkedlist.LinkedList
}

func (stack *LinkedStack) Push(value interface{}) bool {
	return stack.list.Push(value)
}

func (stack *LinkedStack) Pop() (value interface{}, err error) {
	if stack.list.Length == 0 {
		return nil, errors.New("stack is empty")
	}
	return stack.list.Pop(), nil
}

func (stack *LinkedStack) Length() (value int) {
	return stack.list.Length
}
