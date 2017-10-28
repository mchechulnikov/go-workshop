package segment_tree

type node struct {
	sum int
	startIndex int
	endIndex int
	left *node
	right *node
}