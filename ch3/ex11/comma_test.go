package main

import "testing"

func TestComma(t *testing.T) {
	ts := []struct {
		expected string
		s        string
	}{
		{
			s:        "123456",
			expected: "123,456",
		},
		{
			s:        "1",
			expected: "1",
		},
		{
			s:        "1234",
			expected: "1,234",
		},
	}
	for _, tc := range ts {
		got := comma(tc.s)
		if got != tc.expected {
			t.Errorf("got: %v, expected: %v", got, tc.expected)
		}
	}
}
