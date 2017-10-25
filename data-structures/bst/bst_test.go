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

func TestAdd(t *testing.T) {
	values := []int {3, 4, 5, 6, 11, 12, 16, 17}
	bst := New(values)

	t.Run("is contains added", func(t *testing.T) {
		if !bst.Add(10) {
			t.Fatal("Should be added")
		}
		if !bst.Contains(10) {
			t.Fatal("Should be contains")
		}
	})

	t.Run("not added if exists", func(t *testing.T) {
		if bst.Add(11) {
			t.Fatal("Should not be added")
		}
	})

	t.Run("not added twice", func(t *testing.T) {
		if !bst.Add(42) {
			t.Fatal("Should be added")
		}
		if bst.Add(42) {
			t.Fatal("Should not be added")
		}
	})
}

func TestRemove(t *testing.T) {
	values := []int {3, 4, 5, 6, 11, 12, 16, 17}
	bst := New(values)

	t.Run("is not contains removed", func(t *testing.T) {
		if !bst.Remove(11) {
			t.Fatal("Should be removed")
		}
		if bst.Contains(11) {
			t.Fatal("Should be contains")
		}
	})
}