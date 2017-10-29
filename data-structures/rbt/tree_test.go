package rbt

import "testing"

func TestTree_Add(t *testing.T) {
	tree := &Tree{}

	t.Run("", func(t *testing.T) {
		tree.Add(42)
		if !tree.Contains(42) {
			t.Fatal("Not added")
		}
	})
}