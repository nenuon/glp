package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImages(os.Args[1]))
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	if n.Type == html.TextNode {
		words += wordCount(n.Data)
	}
	w, img := countWordsAndImages(n.FirstChild)
	words, images = words+w, images+img
	w, img = countWordsAndImages(n.NextSibling)
	words, images = words+w, images+img
	return
}

func wordCount(str string) int {
	input := bufio.NewScanner(strings.NewReader(str))
	input.Split(bufio.ScanWords)
	result := 0
	for input.Scan() {
		result++
	}
	return result
}
