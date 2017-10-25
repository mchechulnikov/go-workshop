package bfs

import "fmt"

func NewNode(value int) *Node {
	return &Node{value: value}
}

type Node struct {
	value    int
	children []*Node
}

func (n *Node) AddChild(value int) *Node {
	if n.children == nil {
		n.children = make([]*Node, 0)
	}
	node := NewNode(value)
	n.children = append(n.children, node)
	return node
}

func (n *Node) GetChildren() []*Node {
	return n.children
}

func (n *Node) IsChildrenContains(value int) bool {
	if n.value == value {
		return true
	}
	if n.children == nil || len(n.children) <= 0 {
		return false
	}

	tier := make([]*Node, len(n.children))
	copy(tier, n.children)
	for {
		node := tier[0]
		fmt.Println(node.value)
		if node.value == value {
			return true
		}
		if node.children != nil && len(node.children) > 0 {
			tier = append(node.children, tier[1:]...)
		} else if len(tier) > 1 {
			tier = tier[1:]
		} else {
			break
		}
	}
	return false
}
