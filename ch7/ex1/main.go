package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var wc WordCounter
	text := "hello world.\nhoge huga\nI am nenuon.\nこんにちは　世界\n"
	fmt.Fprintf(&wc, "%s", text)
	fmt.Println(wc)
}

type WordCounter struct {
	words, lines int
}

func (w *WordCounter) Write(p []byte) (n int, err error) {
	sc := bufio.NewScanner(strings.NewReader(string(p)))
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		w.words++
	}
	sc = bufio.NewScanner(strings.NewReader(string(p)))
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		w.lines++
	}
	return len(p), nil
}
