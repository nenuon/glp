package main

import (
	"fmt"
)

func main() {
	b := []byte("       a\tb\rc\vd")
	fmt.Printf("%s\n", b)
	fmt.Printf("%s\n", compress(b))
}

func compress(bytes []byte) []byte {
	r := bytes[:0]
	for i := 0; i < len(bytes); i++ {
		if isSpace(bytes[i]) {
			if i == 0 || r[len(r)-1] != ' ' {
				r = append(r, ' ')
			}
			continue
		}
		r = append(r, bytes[i])
	}
	return r
}

func isSpace(b byte) bool {
	switch b {
	case '\t':
		return true
	case '\n':
		return true
	case '\v':
		return true
	case '\f':
		return true
	case '\r':
		return true
	case ' ':
		return true
	case 0x85:
		return true
	case 0xa0:
		return true
	default:
		return false
	}
}
