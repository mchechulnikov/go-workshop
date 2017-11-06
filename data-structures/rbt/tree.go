package rbt

type Tree struct {
	root *node
}

func (t *Tree) Add(value int) bool {
	if t.root == nil {
		t.root = &node{value: value}
		return true
	}

	var currentNode *node
	pointer := t.root
	for {
		if pointer.value < value {
			if pointer.right == nil {
				currentNode = &node {value: value, isRed: true, parent: pointer}
				pointer.right = currentNode
				break
			} else {
				pointer = pointer.right
			}
		} else if pointer.value > value {
			if pointer.left == nil {
				currentNode = &node {value: value, isRed: true, parent: pointer}
				pointer.left = currentNode
				break
			} else {
				pointer = pointer.left
			}
		} else {
			return false
		}
	}
	t.balancingAfterAdding(currentNode)
	return true
}

func (t *Tree) Remove(value int) bool {
	node := t.get(value)
	if node == nil {
		return false
	}
	parent := node.parent

	if node.left == nil && node.right == nil {
		if parent.left == node {
			parent.left = nil
		} else {
			parent.right = nil
		}
		node.parent = nil
	} else if node.left == nil {
		if parent.left == node {
			parent.left = node.right
		} else {
			parent.right = node.right
		}
		node.right.parent = parent
		node.parent = nil
	} else if node.right == nil {
		if parent.left == node {
			parent.left = node.left
		} else {
			parent.right = node.left
		}
		node.left.parent = parent
		node.parent = nil
	} else {
		nextNode := node.right.getMostlyLeft()
		node.value = nextNode.value
		nextNodeParent := nextNode.parent
		if nextNode.right == nil {
			nextNodeParent.left = nil
		} else {
			nextNodeParent.left = nextNode.right
			nextNode.right.parent = nextNodeParent.left
		}
		nextNode.parent = nil
	}
	t.balancingAfterRemoving(node)
	return true
}

func (t *Tree) Contains(value int) bool {
	node := t.get(value)
	if node == nil {
		return false
	}
	return true
}

func (t *Tree) get(value int) *node {
	currentNode := t.root
	for {
		if currentNode == nil {
			return nil
		}
		if currentNode.value > value {
			currentNode = currentNode.left
		} else if currentNode.value < value {
			currentNode = currentNode.right
		} else {
			return currentNode
		}
	}
}

func (t *Tree) balancingAfterAdding(node *node) {
	for {
		if node.isRed && node.parent.isRed {
			uncle := node.getUncle()
			if uncle != nil && uncle.isRed {
				// recolor
				node.parent.isRed = false
				uncle.isRed = false
				if node.parent.parent != t.root {
					node.parent.parent.isRed = true
				}
			} else {
				t.rotate(node.parent)
			}
		}
		node = node.parent
		if node == nil || node == t.root {
			return
		}
	}
}

func (t *Tree) rotate(node *node) {
	parent, grandParent := node.parent, node.parent.parent
	if parent == t.root {
		t.root = node
	}
	if grandParent != nil {
		if grandParent.left == parent {
			grandParent.left = node
		} else {
			grandParent.right = node
		}
	}

	if parent.left == node {
		parent.left = node.right
		node.right = parent
	} else {
		parent.right = node.left
		node.left = parent
	}
	node.parent = grandParent
	parent.parent = node
	parent.isRed = true
	node.isRed = false
}

func (t *Tree) balancingAfterRemoving(removedNode *node) {
	if removedNode.isRed {
		return
	}
	node := removedNode.getBrother()
	for {
		if node.isRed {
			t.rotate(node)
		} else {
			if (node.left == nil || !node.left.isRed) && (node.right == nil || !node.right.isRed) {
				// both children of brother is black
				// recolor
				node.isRed = true
				node.parent.isRed = false
			} else if (node.left.isRed) && (node.right == nil || !node.right.isRed) {
				// left child of brother is red, right child is black
				t.rotate(node.left) // ?? maybe it's wrong
			} else {
				// left child of brother is black, right child is red
				t.rotate(node.right) // ?? maybe it's wrong
			}
		}
		node = node.parent
		if node == nil || node == t.root {
			return
		}
	}
}