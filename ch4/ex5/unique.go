package main

import "fmt"

func main() {
	strings := []string{"hoge", "hoge", "huga"}
	fmt.Println(unique(strings))
}

func unique(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			strings[i] = ""
		}
	}
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}
