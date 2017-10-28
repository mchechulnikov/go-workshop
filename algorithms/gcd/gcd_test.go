package gcd

import "testing"

func TestCalc(t *testing.T) {
	t.Run("0 gcd 2 — error", func(t *testing.T) {
		_, err := Calc(0, 2)
		if err == nil {
			t.Fatal("error fould be nil")
		}
	})

	t.Run("2 gcd 0 — error", func(t *testing.T) {
		_, err := Calc(2, 0)
		if err == nil {
			t.Fatal("error fould be nil")
		}
	})

	t.Run("2 gcd 2 == 2", func(t *testing.T) {
		gcd, err := Calc(2, 2)
		if err != nil {
			t.Fatal(err)
		}
		if gcd != 2 {
			t.Fatalf("Expected %d, actual %d", 2, gcd)
		}
	})

	t.Run("9 gcd 6 == 3", func(t *testing.T) {
		gcd, err := Calc(9, 6)
		if err != nil {
			t.Fatal(err)
		}
		if gcd != 3 {
			t.Fatalf("Expected %d, actual %d", 3, gcd)
		}
	})

	t.Run("6 gcd 9 == 3", func(t *testing.T) {
		gcd, err := Calc(6, 9)
		if err != nil {
			t.Fatal(err)
		}
		if gcd != 3 {
			t.Fatalf("Expected %d, actual %d", 3, gcd)
		}
	})

	t.Run("1112 gcd 695 == 139", func(t *testing.T) {
		gcd, err := Calc(1112, 695)
		if err != nil {
			t.Fatal(err)
		}
		if gcd != 139 {
			t.Fatalf("Expected %d, actual %d", 139, gcd)
		}
	})
}
