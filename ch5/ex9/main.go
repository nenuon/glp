package main

import (
	"fmt"
	"strings"
)

func main() {
	words := "hello $world\n"
	fmt.Print(expand(words, strings.ToUpper))
}

// $がついた単語をf()で置換する
// ex. `hello $world` -> `hello f(world)`
func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	result := make([]string, 0, len(words))
	for _, word := range words {
		if strings.HasPrefix(word, "$") {
			result = append(result, f(word[1:]))
		} else {
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}
