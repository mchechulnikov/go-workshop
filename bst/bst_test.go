package bst

import "testing"

func TestContains(t *testing.T) {
	values := []int {3, 4, 5, 6, 11, 12, 16, 17}
	bst := New(values)

	t.Run("is contains any", func(t *testing.T) {
		if !bst.Contains(11) {
			t.Fatal("Should be contains")
		}
	})

	t.Run("is contains tail element", func(t *testing.T) {
		if !bst.Contains(17) {
			t.Fatal("Should be contains")
		}
	})
}