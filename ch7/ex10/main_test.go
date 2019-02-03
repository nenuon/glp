package ex10

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	ts := []struct {
		vals     []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{1, 2, 3, 2, 1}, true},
		{[]int{1, 2, 2, 1}, true},
		{[]int{}, true},
	}
	for _, tc := range ts {
		if got := IsPalindrome(sort.IntSlice(tc.vals)); got != tc.expected {
			t.Errorf("unexpected result %#v. expected %v, but got %v\n", tc.vals, tc.expected, got)
		}
	}
}
