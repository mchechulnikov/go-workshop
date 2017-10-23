package bst

type node struct {
	value int
	left  *node
	right *node
}

func (n *node) searchNode(value int) (node *node , isFounded bool) {
	if n.value < value {
		if n.right == nil {
			return n, false
		}
		return n.right.searchNode(value)
	} else if n.value > value {
		if n.left == nil {
			return n, false
		}
		return n.left.searchNode(value)
	} else {
		return n, true
	}
}

func (n *node) addChild(value int) {
	node := &node{
		value,
		n.left,
		n.right,
	}
	if n.value < value {
		n.right = node
	} else {
		n.left = node
	}
}

func (n *node) removeChild(value int, parent *node) bool {
	if n.value < value {
		if n.right == nil {
			return false
		}
		return n.right.removeChild(value, n)
	} else if n.value > value {
		if n.left == nil {
			return false
		}
		return n.left.removeChild(value, n)
	} else {
		if n.left == nil && n.right == nil {
			if parent.left == n {
				parent.left = nil
			} else {
				parent.right = nil
			}
		} else if n.left == nil {
			n.value = n.right.value
			n.right = nil
		} else if n.right == nil {
			n.value = n.left.value
			n.left = nil
		} else {
			if n.right.left == nil {
				n.value = n.right.value
				n.right = n.right.right
			} else {
				leftmost, leftmostParent := n.getLeftmost()
				n.value = leftmost.value
				leftmostParent.left = nil
			}
		}
		return true
	}
}

func (n *node) getLeftmost() (node *node, parent *node) {
	node = n
	parent = nil
	for node.left != nil {
		node = node.left
		parent = node
	}
	return node, parent
}