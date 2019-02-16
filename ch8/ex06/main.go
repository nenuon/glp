package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Links struct {
	URL []string
	Dep uint
}

type Link struct {
	URL string
	Dep uint
}

func main() {
	depth := flag.Uint("depth", 0, "depth")
	flag.Parse()

	worklist := make(chan Links)
	unseenlink := make(chan Link)

	go func() { worklist <- Links{flag.Args(), 0} }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenlink {
				go func(link Link) {
					findlist := crawl(link.URL)
					worklist <- Links{findlist, link.Dep + 1}
				}(link)
			}
		}()
	}

	seen := make(map[string]bool)
	for work := range worklist {
		if work.Dep > *depth {
			continue
		}
		for _, url := range work.URL {
			if seen[url] {
				continue
			}
			seen[url] = true
			unseenlink <- Link{url, work.Dep}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 不正なurlを無視
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
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
