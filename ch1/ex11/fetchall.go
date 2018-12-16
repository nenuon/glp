package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Open("website.txt")
	counts := 0
	if err != nil {
		fmt.Fprintf(os.Stderr, "file open: %v\n", err)
		os.Exit(1)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		url := input.Text()
		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}
		counts++
		go fetch(url, ch)
	}
	for i := 0; i < counts; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	f, err := os.Create("fetch_log")
	if err != nil {
		fmt.Printf("file create: %v\n", err)
	}
	defer f.Close()
	nbytes, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
