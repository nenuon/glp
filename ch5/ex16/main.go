package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(join("/", "a", "b", "c"))
}

func join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
