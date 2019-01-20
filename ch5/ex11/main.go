package main

import (
	"fmt"
	"os"
)

var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true},
	"data structures": {"discrete math": true},
	"databases":       {"data structures": true},
	"discrete math":   {"intro to programming": true},
	"networks":        {"operating systems": true},
	"operating systems": {
		"data structures":       true,
		"computer organization": true},
}

func main() {
	result, ok := toposort(prereqs)
	if !ok {
		fmt.Println("roop!")
		os.Exit(1)
	}
	for i, item := range result {
		fmt.Printf("%v:\t%v\n", i, item)
	}
}

func toposort(m map[string]map[string]bool) ([]string, bool) {
	const (
		YET = iota
		NOW
		VISITED
	)
	var order []string
	seen := make(map[string]int)
	var ok bool
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			switch seen[item] {
			case YET:
				seen[item] = NOW
				visitAll(m[item])
				order = append(order, item)
				seen[item] = VISITED
			case NOW:
				ok = false
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	visitAll(keys)
	return order, ok
}
