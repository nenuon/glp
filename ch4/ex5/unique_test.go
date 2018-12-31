package main

import "testing"

func TestUnique(t *testing.T) {
	ts := []struct {
		v        []string
		expected []string
	}{
		{
			[]string{"a", "a", "b", "b", "a", "hoge"},
			[]string{"a", "b", "a", "hoge"},
		},
	}
	for _, tc := range ts {
		if got := unique(tc.v); !equalSlice(got, tc.expected) {
			t.Errorf("test failed: expected %v, but got %v\n", tc.expected, got)
		}
	}
}

func equalSlice(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}
