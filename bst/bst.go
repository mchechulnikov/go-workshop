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
	length := endIndex - startIndex
	k := length / 2
	if length < 1 {
		return nil
	}

	left := getMiddleNode(values, startIndex, startIndex+k)
	var right *node = nil
	if startIndex+k+1 == endIndex {
		right = &node {values[endIndex], nil, nil}
	} else {
		right = getMiddleNode(values, startIndex+k+1, endIndex)
	}
	return &node {values[startIndex + k], left, right}
}