package main

import (
	"fmt"
	"strings"
	"testing"
)

// ダメなassert関数
// 意味がわからない
func assertEqual(x, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func TestSplit(t *testing.T) {
	ts := []struct {
		s    string
		sep  string
		want int // splitした要素の数
	}{
		{
			s:    "a:b:c",
			sep:  ":",
			want: 3,
		},
		{
			s:    "a b:c",
			sep:  " ",
			want: 2,
		},
		{
			s:    "あああ/い",
			sep:  "/",
			want: 2,
		},
		{
			s:    "hoge",
			sep:  "",
			want: 4,
		},
	}
	for _, tc := range ts {
		if got := len(strings.Split(tc.s, tc.sep)); got != tc.want {
			t.Errorf("split(%q, %q) returned %d word, want %d", tc.s, tc.sep, got, tc.want)
		}
	}
}
