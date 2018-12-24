package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	s, t := os.Args[1], os.Args[2]
	fmt.Println(anagram(s, t))
}

func anagram(s, t string) bool {
	return reflect.DeepEqual(count(s), count(t))
}

func count(s string) map[byte]int {
	m := make(map[byte]int)
	for _, c := range s {
		m[byte(c)]++
	}
	return m
}
