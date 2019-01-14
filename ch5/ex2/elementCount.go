package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	count := Count(doc)
	fmt.Printf("%v:\t%v\n", "element name", "times")
	for k, v := range count {
		fmt.Printf("%v:\t%v\n", k, v)
	}
}

func Count(doc *html.Node) map[string]int {
	result := make(map[string]int)
	visit(result, doc)
	return result
}

func visit(mp map[string]int, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		mp[n.Data]++
	}
	visit(mp, n.FirstChild)
	visit(mp, n.NextSibling)
}
