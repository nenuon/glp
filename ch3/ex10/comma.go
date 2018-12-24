package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("9283479487239874"))
}

func comma(s string) string {
	var buf bytes.Buffer
	for len(s) > 3 {
		buf.WriteString(s[:3])
		buf.WriteString(",")
		s = s[3:]
	}
	buf.WriteString(s)
	return buf.String()
}
