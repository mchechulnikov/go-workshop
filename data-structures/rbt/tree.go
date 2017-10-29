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
		if node.parent.isRed {
			uncle := node.getUncle()
			if uncle.isRed {
				node.parent.isRed = false
				uncle.isRed = false
				if node.parent.parent == t.root {
					return
				} else {
					node.parent.parent.isRed = true
				}
				node = node.parent
			} else {
				t.pivot(node.parent)
			}
		}
	}
}

func (t *Tree) pivot(node *node) {
	parent, grandParent := node.parent, node.parent.parent
	if grandParent.left == parent {
		grandParent.left = node
	} else {
		grandParent.right = node
	}
	if parent.left == node {
		parent.left = node.right
		node.right = parent
	} else {
		parent.right = node.left
		node.left = parent
	}
	parent.isRed = true
	node.isRed = false
}
