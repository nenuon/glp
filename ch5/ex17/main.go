package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	start := flag.String("s", "https://golang.org/", "start site")
	flag.Parse()

	resp, err := http.Get(*start)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("http status: %v", resp.Status)
		os.Exit(1)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	divs := ElementsByTagName(doc, "div")

	for _, d := range divs {
		fmt.Println(d)
	}

}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	nameMap := make(map[string]bool)
	for _, na := range name {
		nameMap[na] = true
	}
	nodes := []*html.Node{}
	eachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {
			if nameMap[n.Data] {
				nodes = append(nodes, n)
			}
		}
	}, nil)
	return nodes
}

func eachNode(n *html.Node, pre, post func(n *html.Node)) {
	if n == nil {
		return
	}
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		eachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
