package bfs

import "testing"

func TestNode_IsChildrenContains(t *testing.T) {
	graph := NewNode(1)
	child11 := graph.AddChild(11)
	child12 := graph.AddChild(12)
	child13 := graph.AddChild(13)
	child111 := child11.AddChild(111)
	child11.AddChild(112)
	child113 := child11.AddChild(113)
	child12.AddChild(121)
	child111.AddChild(1111)
	child111.AddChild(1112)
	child111.AddChild(1113)
	child111.AddChild(1114)
	child113.AddChild(1131)
	child13.AddChild(131)

	t.Run("existing value", func(t *testing.T) {
		if !graph.IsChildrenContains(1113) {
			t.Fatalf("Not found existing value")
		}
	})

	t.Run("not existing value", func(t *testing.T) {
		if graph.IsChildrenContains(42) {
			t.Fatalf("Found not existing value")
		}
	})
}
