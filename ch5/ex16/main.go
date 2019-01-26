package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("/", "a", "b", "c"))
}

func join(sep string, a ...string) string {
	if len(a) == 0 {
		return ""
	}
	str := []string{}
	for _, s := range a {
		str = append(str, s)
	}
	return strings.Join(str, sep)
}
