package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	ts := []struct {
		words    string
		f        func(string) string
		expected string
	}{
		{
			words:    `hello world`,
			f:        strings.ToUpper,
			expected: `hello world`,
		},
		{
			words:    `hello $world`,
			f:        strings.ToUpper,
			expected: `hello WORLD`,
		},
		{
			words:    `$hello world $hoge $fuga`,
			f:        strings.ToUpper,
			expected: `HELLO world HOGE FUGA`,
		},
	}

	for _, tc := range ts {
		if got := expand(tc.words, tc.f); got != tc.expected {
			t.Errorf("unexpected result. expected '%v', but got '%v'\n", tc.expected, got)
		}
	}
}
