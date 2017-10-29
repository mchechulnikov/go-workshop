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