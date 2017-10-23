package bst

import "sort"

func New(values []int) *Tree {
	copiedValues := getSortedCopy(values)
	root := getMiddleNode(copiedValues, 0, len(copiedValues)-1)
	return &Tree{root}
}

func getSortedCopy(values []int) []int {
	copiedValues := make([]int, len(values))
	copy(copiedValues, values)
	sort.Ints(copiedValues)
	return copiedValues
}

func getMiddleNode(values []int, startIndex int, endIndex int) *node {
	k := (endIndex - startIndex) / 2
	return &node {
		values[k],
		getMiddleNode(values, startIndex, k),
		getMiddleNode(values, k + 1, endIndex),
	}
}