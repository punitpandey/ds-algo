package stack

import "errors"

type Stack struct {
	top  int
	list []interface{}
}

func (stack *Stack) Push(value interface{}) {
	stack.list = append(stack.list, value)
	stack.top++
}

func (stack *Stack) Pop() (value interface{}, err error) {
	if stack.top <= 0 {
		return nil, errors.New("stack is empty")
	}
	value = stack.list[stack.top-1]
	stack.list = stack.list[:stack.top-1]
	stack.top--
	return
}

func (stack *Stack) Length() (value int) {
	return stack.top
}
