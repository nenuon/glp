package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	charCount()
}

func charCount() {
	var mp = make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, nbytes, err
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		if unicode.IsLetter(r) {
			mp["letter"]++
		} else if unicode.IsDigit(r) {
			mp["disit"]++
		} else {
			mp["other"]++
		}
	}
	fmt.Printf("unicode\tcounts\n")
	for k, v := range mp {
		fmt.Printf("%v\t%v\n", k, v)
	}
}
