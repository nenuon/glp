package main

import (
	"fmt"
)

func main() {
	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	rotate2(v, -100000)
	fmt.Println(v)
}

// r > 0右回転
// r < 0左回転
func rotate(v []int, r int) []int {
	z := make([]int, len(v), len(v))
	n := len(v)
	for i := 0; i < n; i++ {
		z[(i+n+r)%n] = v[i]
	}
	return z
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func rotate2(v []int, r int) {
	L := len(v)
	if r = r % L; r < 0 {
		r = L + r
	}
	g := gcd(L, r)
	for i := 0; i < g; i++ {
		x := v[i]
		for j := (i + r + L) % L; j != i; j = (j + r + L) % L {
			x, v[j] = v[j], x
		}
		v[i] = x
	}
}
