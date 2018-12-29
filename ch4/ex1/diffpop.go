package main

import (
	"crypto/sha256"
	"fmt"

	"./pop"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	sum := 0
	for i := 0; i < 32; i++ {
		sum += pop.SamePopCount(uint64(c1[i]), uint64(c2[i]))
	}
	fmt.Println(sum)
}
