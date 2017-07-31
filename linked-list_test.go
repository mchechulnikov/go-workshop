package main

import (
	"testing"
)

func TestLinkedList_Get(t *testing.T) {
	t.Run("empty list — error", func(t *testing.T) {
		list := New()
		_, err := list.Get(0)
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})

	t.Run("not empty list, valid index — valid element", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Add(44)
		list.Add(45)
		list.Add(46)
		element, err := list.Get(3)
		if err != nil {
			t.Fatalf("Error not eq nil: %s", err.Error())
		}
		if element.GetValue() != 45 {
			t.Fatalf("Expected %d, but received %d", 45, element.GetValue())
		}
	})
}

func TestLinkedList_Add(t *testing.T) {
	t.Run("empty list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		if list.Length != 1 {
			t.Fatalf("Invalid length")
		}
		if list.First.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}
	})

	t.Run("not empty list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Add(44)
		if list.Length != 3 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, secondElement.GetValue())
		}

		thirdElement := secondElement.GetNext()
		if thirdElement.GetValue() != 44 {
			t.Fatalf("Expected %d, but received %d", 44, thirdElement.GetValue())
		}
	})
}

func TestLinkedList_AddByIndex(t *testing.T) {
	t.Run("empty list, valid index — success", func(t *testing.T) {
		list := New()
		err := list.AddByIndex(42, 0)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if list.Length != 1 {
			t.Fatalf("Invalid length")
		}
		if list.First.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}
	})

	t.Run("empty list, invalid index — error", func(t *testing.T) {
		list := New()
		err := list.AddByIndex(42, 1)
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})

	t.Run("not empty list, add to begin of the list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.AddByIndex(41, 0)
		if list.Length != 3 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 41 {
			t.Fatalf("Expected %d, but received %d", 41, list.First.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, secondElement.GetValue())
		}

		thirdElement := secondElement.GetNext()
		if thirdElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, thirdElement.GetValue())
		}
	})

	t.Run("not empty list, add between elements — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Add(45)
		list.AddByIndex(44, 2)
		if list.Length != 4 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, secondElement.GetValue())
		}

		thirdElement := secondElement.GetNext()
		if thirdElement.GetValue() != 44 {
			t.Fatalf("Expected %d, but received %d", 44, thirdElement.GetValue())
		}

		fourthElement := thirdElement.GetNext()
		if fourthElement.GetValue() != 45 {
			t.Fatalf("Expected %d, but received %d", 45, fourthElement.GetValue())
		}
	})

	t.Run("not empty list, add to end of the list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.AddByIndex(44, 2)
		if list.Length != 3 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, secondElement.GetValue())
		}

		thirdElement := secondElement.GetNext()
		if thirdElement.GetValue() != 44 {
			t.Fatalf("Expected %d, but received %d", 44, thirdElement.GetValue())
		}
	})
}

func TestLinkedList_Delete(t *testing.T) {
	t.Run("empty list — error", func(t *testing.T) {
		list := New()
		err := list.Delete(0)
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})

	t.Run("not empty list, delete from begin of the list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Delete(0)
		if list.Length != 1 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, list.First.GetValue())
		}
	})

	t.Run("not empty list, delete between elements — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Add(44)
		list.Delete(1)
		if list.Length != 2 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, firstElement.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 44 {
			t.Fatalf("Expected %d, but received %d", 44, secondElement.GetValue())
		}
	})

	t.Run("not empty list, delete from end of the list — success", func(t *testing.T) {
		list := New()
		list.Add(42)
		list.Add(43)
		list.Add(44)
		list.Delete(2)
		if list.Length != 2 {
			t.Fatalf("Invalid length")
		}

		firstElement := list.First
		if firstElement.GetValue() != 42 {
			t.Fatalf("Expected %d, but received %d", 42, list.First.GetValue())
		}

		secondElement := firstElement.GetNext()
		if secondElement.GetValue() != 43 {
			t.Fatalf("Expected %d, but received %d", 43, secondElement.GetValue())
		}

		if secondElement.next != nil {
			t.Fatalf("Third element detected")
		}
	})
}
