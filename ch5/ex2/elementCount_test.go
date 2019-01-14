package main

import (
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func CountTest(b *testing.T) {
	ts := []struct {
		doc      string
		expected map[string]int
	}{
		{
			``,
			map[string]int{},
		},
		{
			`<html><body></html></body>`,
			map[string]int{"html": 1, "body": 1},
		},
		{
			`<html>
			<body>
			<a href="https://google.com">
			<a href="https://yahoo.com">
			<div>
				<a href="https://hithub.com">
			</div>
			</body>
			</html>`,
			map[string]int{"html": 1, "body": 1, "a": 3, "div": 1},
		},
	}
	for _, tc := range ts {
		HTMLdoc, err := html.Parse(strings.NewReader(tc.doc))
		if err != nil {
			fmt.Errorf("%v\n", err)
		}
		if got := Count(HTMLdoc); !equals(got, tc.expected) {
			fmt.Errorf("unexpected count. expected %v, but got %v\n", tc.expected, got)
		}
	}
}
func equals(actual, expect map[string]int) bool {
	if len(actual) != len(expect) {
		return false
	}
	for k, v := range expect {
		if v != actual[k] {
			return false
		}
	}
	return true
}
