package rbt

type node struct {
	value int
	isRed bool
	parent *node
	left *node
	right *node
}

func (n *node) getUncle() *node {
	parent := n.parent
	if parent.parent.left == parent {
		return parent.parent.right
	} else {
		return parent.parent.left
	}
}

func (n *node) getBrother() *node {
	parent := n.parent
	if parent.left == n {
		return parent.right
	} else {
		return parent.left
	}
}

func (n *node) getMostlyLeft() *node {
	node := n
	for {
		if node.left == nil {
			break
		}
		node = node.left
	}
	return node
}