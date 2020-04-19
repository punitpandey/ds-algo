package bst

import "../queue"

type node struct {
	value int
	left  *node
	right *node
}

type BST struct {
	root   *node
	Length int
}

func (tree *BST) Add(value int) {
	var currentNode, parentNode *node
	if tree.root == nil {
		tree.root = &node{
			value: value,
			left:  nil,
			right: nil,
		}
		tree.Length++
	} else {
		currentNode = tree.root
		for currentNode != nil {
			parentNode = currentNode
			if currentNode.value > value {
				currentNode = currentNode.left
			} else {
				currentNode = currentNode.right
			}
		}
		if parentNode.value < value {
			parentNode.right = &node{
				value: value,
				left:  nil,
				right: nil,
			}
		} else {
			parentNode.left = &node{
				value: value,
				left:  nil,
				right: nil,
			}
		}
	}
}

func (tree *BST) Remove(value int) {
	var replacer, replacerParent, parentNode *node
	currentNode := tree.root
	for currentNode != nil {
		if currentNode.value == value {
			if currentNode.right == nil {
				if parentNode.value > currentNode.value {
					parentNode.left = nil
				} else {
					parentNode.right = nil
				}
			} else {
				replacer, replacerParent = currentNode.right.Min()
				if replacerParent != nil {
					replacerParent.left = replacer.right
				}
				replacer.left = currentNode.left
				replacer.right = currentNode.right
				if currentNode != tree.root {
					if parentNode.value > currentNode.value {
						parentNode.left = replacer
					} else {
						parentNode.right = replacer
					}
				} else {
					tree.root = replacer
				}
			}
			return
		} else if currentNode.value < value {
			parentNode = currentNode
			currentNode = currentNode.right
		} else {
			parentNode = currentNode
			currentNode = currentNode.left
		}
	}
}

func (node *node) Min() (minNode, parentNode *node) {
	minNode = node
	for minNode.left != nil {
		parentNode = minNode
		minNode = minNode.left
	}
	return
}

func (tree *BST) PreOrderTraverse() (values []int) {
	return tree.root.PreOrderTraverse()
}

func (node *node) PreOrderTraverse() (values []int) {
	currentNode := node
	values = append(values, currentNode.value)
	if currentNode.left != nil {
		values = append(values, currentNode.left.PreOrderTraverse()...)
	}
	if currentNode.right != nil {
		values = append(values, currentNode.right.PreOrderTraverse()...)
	}
	return
}

func (tree *BST) InOrderTraverse() (value []int) {
	return tree.root.InOrderTraverse()
}

func (node *node) InOrderTraverse() (values []int) {
	currentNode := node
	if currentNode.left != nil {
		values = append(values, currentNode.left.InOrderTraverse()...)
	}
	values = append(values, currentNode.value)
	if currentNode.right != nil {
		values = append(values, currentNode.right.InOrderTraverse()...)
	}
	return
}

func (tree *BST) PostOrderTraverse() (value []int) {
	return tree.root.PostOrderTraverse()
}

func (node *node) PostOrderTraverse() (values []int) {
	currentNode := node
	if currentNode.left != nil {
		values = append(values, currentNode.left.PostOrderTraverse()...)
	}
	if currentNode.right != nil {
		values = append(values, currentNode.right.PostOrderTraverse()...)
	}
	values = append(values, currentNode.value)
	return
}

func (tree *BST) BFTraverse() (values []int) {
	var queue = new(queue.LinkedQueue)
	queue.Enqueue(tree.root)
	for queue.Length() != 0 {
		firstNode, _ := queue.Dequeue()
		currentNode := firstNode.(*node)
		values = append(values, currentNode.value)
		if currentNode.left != nil {
			queue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			queue.Enqueue(currentNode.right)
		}
	}
	return values
}

func (tree *BST) BFSearch(key int) bool {
	var queue = new(queue.LinkedQueue)
	queue.Enqueue(tree.root)
	for queue.Length() != 0 {
		firstNode, _ := queue.Dequeue()
		currentNode := firstNode.(*node)
		if currentNode.value == key {
			return true
		}
		if currentNode.left != nil {
			queue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			queue.Enqueue(currentNode.right)
		}
	}
	return false
}
