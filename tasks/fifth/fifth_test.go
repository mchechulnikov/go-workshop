package fifth

import "testing"

func TestGetSliceWithMaxSum(t *testing.T) {
	values := []int {1, -300, 2, 1, 301, -42, 10}
	result := GetSliceWithMaxSum(values)

	t.Run("expected length", func(t *testing.T) {
		if len(result) != 3 {
			t.Fatalf("Expected %s, but in actual %s", 3, len(result))
		}
	})

	t.Run("expected values", func(t *testing.T) {
		if result[0] != 2 {
			t.Fatalf("Expected %s, but in actual %s", 2, result[0])
		}
		if result[1] != 1 {
			t.Fatalf("Expected %s, but in actual %s", 1, result[1])
		}
		if result[2] != 301 {
			t.Fatalf("Expected %s, but in actual %s", 301, result[2])
		}
	})
}