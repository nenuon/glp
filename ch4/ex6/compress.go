package main

import "fmt"

func main() {
	b := []byte("a       b         c        d")
	fmt.Printf("%s\n", compress(b))
}

func compress(bytes []byte) []byte {
	r := bytes[:1]
	for i := 1; i < len(bytes); i++ {
		if bytes[i] == ' ' && r[len(r)-1] == ' ' {
			continue
		}
		r = append(r, bytes[i])
	}
	return r
}
