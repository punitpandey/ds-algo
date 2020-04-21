package avl

type Node struct {
	value int
	left  *Node
	right *Node
}

type AVL struct {
	root *Node
}

func (avl *AVL) Insert(value int) {
	if avl.root == nil {
		avl.root = &Node{
			value: value,
		}
	} else {
		avl.root = avl.root.Insert(value)
		avl.root = avl.root.Balance()
	}
}

/*
Successor returns successor node and parent of successor as:
Case 1:returns nil, nil if callee has no right child
Case 2:returns node.right, node if successor is right node
Case 3:returns successorNode, successorParent if successor is right node's child
*/
func (node *Node) Successor() (successor, parent *Node) {
	parent = node
	successor = node.right
	for successor != nil && successor.left != nil {
		if successor.left == nil {
			break
		}
		parent = successor
		successor = successor.left
	}
	return successor, parent
}

/*
Remove of *AVL handle head removal cases, or calls Remove of *Node
Case 1: right node is successor
Case 2: right's child node is successor
Case 3: left node is successor
*/
func (avl *AVL) Remove(value int) {
	if avl.root.value == value {
		replacer, replacerParent := avl.root.Successor()
		if replacer != nil {
			replacer.left = avl.root.left
			if replacerParent != avl.root {
				replacerParent.left = replacer.right
				replacer.right = avl.root.right
			}
			avl.root = replacer
			avl.root = avl.root.Balance()
		} else {
			avl.root = avl.root.left
		}
	} else {
		avl.root = avl.root.Remove(value, nil)
	}
}

/*
Remove of *Node handle cases with non-head node removal
Case 1: having only right node
Case 2: having successor other than right node with no child
Case 3: having successor with right child
*/
func (node *Node) Remove(value int, parentNode *Node) *Node {
	if node.value == value {
		if node.right == nil {
			if parentNode.value > node.value {
				parentNode.left = nil
			} else {
				parentNode.right = nil
			}
		} else {
			replacer, replacerParent := node.Successor()
			replacer.left = node.left
			if replacerParent != node {
				replacer.right = node.right
				replacerParent.left = replacer.right
			}
			if parentNode.value > node.value {
				parentNode.left = replacer
			} else {
				parentNode.right = replacer
			}
		}
	} else if node.value < value {
		parentNode = node
		_ = node.right.Remove(value, node)
		parentNode = parentNode.Balance()
	} else {
		parentNode = node
		_ = node.left.Remove(value, node)
		parentNode = parentNode.Balance()
	}
	return parentNode
}

func (avl *AVL) Traverse() []int {
	return avl.root.Traverse()
}

// Traverse returns []int with Pre-Order traversal
func (node *Node) Traverse() (nodes []int) {
	if node == nil {
		return
	}
	nodes = append(nodes, node.value)
	nodes = append(nodes, node.left.Traverse()...)
	nodes = append(nodes, node.right.Traverse()...)
	return nodes
}

// Balance returns the rearrange node with maximum depth difference of 1 for callee
func (node *Node) Balance() *Node {
	height := node.left.Height() - node.right.Height()
	if height < -1 {
		// left rotation required
		diff := node.right.left.Height() - node.right.right.Height()
		if diff > 0 {
			node.right = node.right.RightRotate()
		}
		node = node.LeftRotate()
	}
	if height > 1 {
		// right rotation required
		diff := node.left.left.Height() - node.left.right.Height()
		if diff < 0 {
			node.left = node.left.LeftRotate()
		}
		node = node.RightRotate()
	}
	return node
}

// Height returns the maximum depth of callee node
func (node *Node) Height() int {
	var leftHeight, rightHeight int
	if node == nil {
		return 0
	}
	leftHeight = node.left.Height()
	rightHeight = node.right.Height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func (node *Node) LeftRotate() *Node {
	node.right, node.right.left, node = node.right.left, node, node.right
	return node
}

func (node *Node) RightRotate() *Node {
	node.left, node.left.right, node = node.left.right, node, node.left
	return node
}

func (node *Node) Insert(value int) *Node {
	if value < node.value {
		if node.left != nil {
			node.left = node.left.Insert(value)
			node = node.Balance()
		} else {
			node.left = &Node{
				value: value,
			}
		}
	} else {
		if node.right != nil {
			node.right = node.right.Insert(value)
			node = node.Balance()
		} else {
			node.right = &Node{
				value: value,
			}
		}
	}
	return node
}
