package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	dup(os.Stdout, os.Args)
}

func countLines(f *os.File, counts map[string]int, lineToFilename map[string]map[string]struct{}) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "EOF" {
			break
		}
		counts[input.Text()]++
		if len(lineToFilename[input.Text()]) == 0 {
			lineToFilename[input.Text()] = map[string]struct{}{f.Name(): {}}
		}
		lineToFilename[input.Text()][f.Name()] = struct{}{}
	}
}

func dup(w io.Writer, args []string) {
	counts := make(map[string]int)
	lineToFilenames := make(map[string]map[string]struct{})
	files := args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFilenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "wx4: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFilenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			filename := ""
			sep := ""
			for names := range lineToFilenames[line] {
				filename += sep + names
				sep = " "
			}
			fmt.Fprintf(w, "%d\t%s\t%s\n", n, line, filename)
		}
	}
}
