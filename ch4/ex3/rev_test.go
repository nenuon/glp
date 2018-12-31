package main

import "testing"

func TestReverse(t *testing.T) {
	ts := []struct {
		v        [8]int
		expected [8]int
	}{
		{
			[8]int{1, 2, 3, 4, 5, 6, 7, 8},
			[8]int{8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tc := range ts {
		if got := reverse(&tc.v); *got != tc.expected {
			t.Errorf("test failed: expected %v, but got %v\n", tc.expected, got)
		}
	}
}
