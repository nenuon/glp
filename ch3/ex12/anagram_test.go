package main

import "testing"

func TestAnagram(t *testing.T) {
	ts := []struct {
		s        string
		t        string
		expected bool
	}{
		{"a", "a", true},
		{"a", "b", false},
		{"abc", "abc", true},
		{"cba", "cba", true},
		{"a", "aa", false},
		{"hogehuga", "hugahoge", true},
	}
	for _, tc := range ts {
		if got := anagram(tc.s, tc.t); got != tc.expected {
			t.Errorf("expected %v, but got %v\n", tc.expected, got)
		}
	}
}
