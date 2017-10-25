package hashMap

import (
	"testing"
)

func TestHashMap_Add(t *testing.T) {
	t.Run("empty map, a few values added — success", func(t *testing.T) {
		hmap := New()
		err := hmap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hmap.Length() != 1 {
			t.Fatalf("Invalid length")
		}
		err = hmap.Add("bar", 43)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hmap.Length() != 2 {
			t.Fatalf("Invalid length")
		}
	})

	t.Run("add already existing key — error", func(t *testing.T) {
		hmap := New()
		err := hmap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hmap.Length() != 1 {
			t.Fatalf("Invalid length")
		}
		err = hmap.Add("foo", 43)
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})
}

func TestHashMap_Get(t *testing.T) {
	t.Run("get not presented key — error", func(t *testing.T) {
		hmap := New()
		_, err := hmap.Get("foo")
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})

	t.Run("by existing key — success", func(t *testing.T) {
		hmap := New()
		err := hmap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		value, err := hmap.Get("foo")
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if value != 42 {
			t.Fatalf("Expected %d, but received %d", 42, value)
		}
	})
}

func TestHashMap_Delete(t *testing.T) {
	t.Run("delete key — success", func(t *testing.T) {
		hmap := New()
		hmap.Add("foo", 42)
		hmap.Add("bar", 43)
		hmap.Add("buz", 44)
		if hmap.Length() != 3 {
			t.Fatalf("Expected %d elements in hash map, but actual %d", 3, hmap.Length())
		}

		err := hmap.Delete("bar")
		if err != nil {
			t.Fatalf("Error not eq nil: %s", err.Error())
		}
		if hmap.Length() != 2 {
			t.Fatalf("Expected %d elements in hash map, but actual %d", 2, hmap.Length())
		}
	})

	t.Run("delete not existing key — error", func(t *testing.T) {
		hmap := New()
		err := hmap.Delete("foo")
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})
}
