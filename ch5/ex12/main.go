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
	printHTML(doc)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n == nil {
		return
	}
	if pre != nil {
		pre(n)
	}
	forEachNode(n.FirstChild, pre, post)
	if post != nil {
		post(n)
	}
	forEachNode(n.NextSibling, pre, post)
}

func printHTML(n *html.Node) {
	var depth int
	var startElement func(n *html.Node)
	var endElement func(n *html.Node)

	startElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, " ", n.Data)
			depth++
		}
	}
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, " ", n.Data)
		}
	}

	forEachNode(n, startElement, endElement)
}
