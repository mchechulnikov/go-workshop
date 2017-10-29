package rbt

type Stack struct {
	top *stackNode
}

type stackNode struct {
	value int
	prev *stackNode
}

func (s *Stack) Push(value int) {
	s.top = &stackNode{value: value, prev: s.top}
}

func (s *Stack) Pop() int {
	node := s.top
	s.top = s.top.prev
	return node.value
}