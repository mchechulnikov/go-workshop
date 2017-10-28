package gcd

import "fmt"

func Calc(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("one of parameters eq zero")
	}
	if a == b {
		return a, nil
	}
	if a < b {
		c := a
		a = b
		b = c
	}
	d, gcd := a, b
	for {
		t := d % gcd
		if t == 0 {
			return gcd, nil
		}
		d = gcd
		gcd = t
	}
}
