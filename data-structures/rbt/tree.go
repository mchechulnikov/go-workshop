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
	// todo
	return false
}

func (t *Tree) Contains(value int) bool {
	currentNode := t.root
	for {
		if currentNode == nil {
			return false
		}
		if currentNode.value > value {
			currentNode = currentNode.left
		} else if currentNode.value < value {
			currentNode = currentNode.right
		} else {
			return true
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
