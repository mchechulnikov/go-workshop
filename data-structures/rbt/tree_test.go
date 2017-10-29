package rbt

import "testing"

func TestTree_Add(t *testing.T) {
	tree := &Tree{}

	t.Run("add one", func(t *testing.T) {
		tree.Add(42)
		if !tree.Contains(42) {
			t.Fatal("Not added")
		}
	})

	t.Run("add a few", func(t *testing.T) {
		tree.Add(42)
		tree.Add(41)
		tree.Add(43)
		if !tree.Contains(42) {
			t.Fatal("Not added")
		}
		if !tree.Contains(41) {
			t.Fatal("Not added")
		}
		if !tree.Contains(42) {
			t.Fatal("Not added")
		}
	})

	t.Run("adds with rotate and recolor", func(t *testing.T) {
		tree.Add(42)
		tree.Add(41)
		tree.Add(43)

		// add with one recolor
		tree.Add(44)
		tree.Add(46)
		if !tree.Contains(44) {
			t.Fatal("Not added")
		}
		if !tree.Contains(46) {
			t.Fatal("Not added")
		}
		consideredNode := tree.root.right
		if consideredNode.value != 44 || consideredNode.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.left.value != 43 || !consideredNode.left.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.right.value != 46 || !consideredNode.right.isRed {
			t.Fatal("Not expected node")
		}

		// add with one rotate
		tree.Add(47)
		consideredNode = tree.root.right
		if consideredNode.value != 44 || !consideredNode.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.left.value != 43 || consideredNode.left.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.right.value != 46 || consideredNode.right.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.right.right.value != 47 || !consideredNode.right.right.isRed {
			t.Fatal("Not expected node")
		}

		// add with another one recolor
		tree.Add(48)
		consideredNode = tree.root.right.right
		if consideredNode.value != 47 || consideredNode.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.left.value != 46 || !consideredNode.left.isRed {
			t.Fatal("Not expected node")
		}
		if consideredNode.right.value != 48 || !consideredNode.right.isRed {
			t.Fatal("Not expected node")
		}

		// add with recolor and then rotate
		tree.Add(45)
		if tree.root.value != 44 || tree.root.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.left.value != 42 || !tree.root.left.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.left.left.value != 41 || tree.root.left.left.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.left.right.value != 43 || tree.root.left.right.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.right.value != 47 || !tree.root.right.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.right.left.value != 46 || tree.root.right.left.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.right.right.value != 48 || tree.root.right.right.isRed {
			t.Fatal("Not expected node")
		}
		if tree.root.right.left.left.value != 45 || !tree.root.right.left.left.isRed {
			t.Fatal("Not expected node")
		}
	})
}