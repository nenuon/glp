package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	var buf bytes.Buffer
	for len(s) > 3 {
		size := len(s) % 3
		if size == 0 {
			size = 3
		}
		buf.WriteString(s[:size])
		buf.WriteString(",")
		s = s[size:]
	}
	buf.WriteString(s)
	return buf.String()
}
