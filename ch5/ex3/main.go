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
		log.Fatalln(err)
	}
	printDocument(doc)
}

func printDocument(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.TextNode {
		fmt.Print(n.Data)
	}
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		printDocument(n.FirstChild)
	}
	printDocument(n.NextSibling)
}
