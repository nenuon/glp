package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(comma(os.Args[1]))
}

func comma(s string) string {
	var buf bytes.Buffer
	ss := s
	if n := strings.Index(s, "."); n >= 0 {
		ss = s[:n]
	}
	for len(ss) > 3 {
		size := len(ss) % 3
		if size == 0 {
			size = 3
		}
		if strings.Contains(ss[:size], ".") {
			break
		}
		buf.WriteString(ss[:size])
		buf.WriteString(",")
		ss = ss[size:]
		s = s[size:]
	}
	buf.WriteString(s)
	return buf.String()
}
