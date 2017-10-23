package bst

type Tree struct {
	root *node
}

func (t *Tree) Contains(value int) bool {
	_, isFounded := t.root.searchNode(value)
	return isFounded
}

func (t *Tree) Add(value int) bool {
	node, isFounded := t.root.searchNode(value)
	if isFounded {
		return false
	} else {
		node.addChild(value)
		return true
	}
}

func (t *Tree) Remove(value int) bool {
	node, isFounded := t.root.searchNode(value)
	if isFounded {
		node.addChild(value) // todo
		return true
	} else {
		return false
	}
}