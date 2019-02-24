package main

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	ts := []struct {
		in   string
		want map[rune]int
	}{
		{
			in:   "a",
			want: map[rune]int{'a': 1},
		},
		{
			in:   "",
			want: map[rune]int{},
		},
		{
			in:   "hoge huga",
			want: map[rune]int{'h': 2, 'o': 1, 'g': 2, 'e': 1, 'a': 1, 'u': 1, ' ': 1},
		},
		{
			in:   "ほげふが",
			want: map[rune]int{'ほ': 1, 'げ': 1, 'ふ': 1, 'が': 1},
		},
	}
	for _, tc := range ts {
		if got := charcount(bytes.NewReader([]byte(tc.in))); !isSame(tc.want, got) {
			t.Errorf("unexpected counts: expected %v, but got %v\n", tc.want, got)
		}
	}
}

func isSame(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
