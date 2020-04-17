package linkedlist

type node struct {
	value interface{}
	next  *node
}
type LinkedList struct {
	root   *node
	Length int
}

func (list *LinkedList) Insert(value interface{}, position int) bool {
	if position > list.Length {
		return false
	} else if position == 0 {
		newNode := &node{
			value: value,
			next:  list.root,
		}
		list.root = newNode
		list.Length++
	} else {
		currentNode := list.root
		currentIndex := 0
		for ; currentIndex < position-1; currentIndex++ {
			currentNode = currentNode.next
		}
		newNode := &node{
			value: value,
			next:  currentNode.next,
		}
		currentNode.next = newNode
		list.Length++
	}
	return true
}

func (list *LinkedList) Push(value interface{}) bool {
	return list.Insert(value, list.Length)
}

func (list *LinkedList) Unshift(value interface{}) bool {
	return list.Insert(value, 0)
}

func (list *LinkedList) Remove(position int) (value interface{}) {
	if list.Length == 0 || position >= list.Length {
		return nil
	}
	if position == 0 {
		value = list.root.value
		list.root = list.root.next
		list.Length--
	} else {
		currentNode := list.root
		currentIndex := 0
		for ; currentIndex < position-1; currentIndex++ {
			currentNode = currentNode.next
		}
		value = currentNode.next.value
		currentNode.next = currentNode.next.next
		list.Length--
	}
	return
}

func (list *LinkedList) Pop() interface{} {
	return list.Remove(list.Length - 1)
}

func (list *LinkedList) Shift() interface{} {
	return list.Remove(0)
}

func (list *LinkedList) Traverse() (response []interface{}) {
	currentNode := list.root
	for currentNode != nil {
		response = append(response, currentNode.value)
		currentNode = currentNode.next
	}
	return
}
