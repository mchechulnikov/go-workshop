package main

import (
	"testing"
)

func TestHashMap_Add(t *testing.T) {
	t.Run("empty map, a few values added — success", func(t *testing.T) {
		hashMap := NewHashMap()
		err := hashMap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hashMap.Length() != 1 {
			t.Fatalf("Invalid length")
		}
		err = hashMap.Add("bar", 43)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hashMap.Length() != 2 {
			t.Fatalf("Invalid length")
		}
	})

	t.Run("add already existed key — error", func(t *testing.T) {
		hashMap := NewHashMap()
		err := hashMap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if hashMap.Length() != 1 {
			t.Fatalf("Invalid length")
		}
		err = hashMap.Add("foo", 43)
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})
}

func TestHashMap_Get(t *testing.T) {
	t.Run("get not presented key — error", func(t *testing.T) {
		hashMap := NewHashMap()
		_, err := hashMap.Get("foo")
		if err == nil {
			t.Fatalf("Error eq nil")
		}
	})

	t.Run("by existed key — success", func(t *testing.T) {
		hashMap := NewHashMap()
		err := hashMap.Add("foo", 42)
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		value, err := hashMap.Get("foo")
		if err != nil {
			t.Fatalf("Error not eq nil")
		}
		if value != 42 {
			t.Fatalf("Expected %d, but received %d", 42, value)
		}
	})
}
