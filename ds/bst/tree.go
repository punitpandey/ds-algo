package bst

import "../queue"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root   *Node
	Length int
}

func (tree *BST) Add(value int) {
	var currentNode, parentNode *Node
	if tree.root == nil {
		tree.root = &Node{
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
			parentNode.right = &Node{
				value: value,
				left:  nil,
				right: nil,
			}
		} else {
			parentNode.left = &Node{
				value: value,
				left:  nil,
				right: nil,
			}
		}
	}
}

func (tree *BST) Remove(value int) {
	var replacer, replacerParent, parentNode *Node
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

func (node *Node) Min() (minNode, parentNode *Node) {
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

func (node *Node) PreOrderTraverse() (values []int) {
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

func (node *Node) InOrderTraverse() (values []int) {
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

func (node *Node) PostOrderTraverse() (values []int) {
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

func (tree *BST) BFTraverse() (value []int) {
	var queue = new(queue.LinkedQueue)
	queue.Enqueue(tree.root)
	for queue.Length() != 0 {
		node, _ := queue.Dequeue()
		currentNode := node.(*Node)
		value = append(value, currentNode.value)
		if currentNode.left != nil {
			queue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			queue.Enqueue(currentNode.right)
		}
	}
	return value
}
