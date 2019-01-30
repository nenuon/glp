package main

import "testing"

func newIntSet(set ...int) *IntSet {
	ret := &IntSet{}
	for _, s := range set {
		ret.Add(s)
	}
	return ret
}

func TestAddAll(t *testing.T) {
	ts := []struct {
		set      *IntSet
		add      []int
		expected string
	}{
		{
			set:      newIntSet(1, 3, 5),
			add:      []int{2, 4},
			expected: `{1 2 3 4 5}`,
		},
	}
	for _, tc := range ts {
		tc.set.AddAll(tc.add...)
		if got := tc.set.String(); got != tc.expected {
			t.Errorf("unexpected result: expected %v, but got %v\n", tc.expected, got)
		}
	}
}

func TestLen(t *testing.T) {
	ts := []struct {
		set      *IntSet
		expected int
	}{
		{
			set:      &IntSet{words: []uint64{11}}, // {0 1 3}
			expected: 3,
		},
		{
			set:      &IntSet{words: []uint64{}}, // {}
			expected: 0,
		},
	}
	for _, tc := range ts {
		if got := tc.set.Len(); got != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, got)
		}
	}
}

func TestRemove(t *testing.T) {
	ts := []struct {
		set      *IntSet
		del      int
		expected string
	}{
		{
			set:      &IntSet{words: []uint64{11}}, // {0 1 3}
			del:      1,
			expected: `{0 3}`,
		},
	}
	for _, tc := range ts {
		if tc.set.Remove(tc.del); tc.set.String() != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, tc.set.String())
		}
	}
}

func TestClear(t *testing.T) {
	ts := []struct {
		set      *IntSet
		expected string
	}{
		{
			set:      &IntSet{words: []uint64{11}}, // {0 1 3}
			expected: `{}`,
		},
	}
	for _, tc := range ts {
		if tc.set.Clear(); tc.set.String() != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, tc.set.String())
		}
	}
}

func TestCopy(t *testing.T) {
	ts := []struct {
		set      *IntSet
		expected string
	}{
		{
			set:      &IntSet{words: []uint64{11}}, // {0 1 3}
			expected: `{0 1 3}`,
		},
	}
	for _, tc := range ts {
		if got := tc.set.Copy(); got.String() != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, got.String())
		}
	}
}

func TestIntersectWith(t *testing.T) {
	ts := []struct {
		s        *IntSet
		t        *IntSet
		expected string
	}{
		{
			s:        newIntSet(1, 3, 5, 7),
			t:        newIntSet(4, 5, 7),
			expected: `{5 7}`,
		},
		{
			s:        newIntSet(1, 3, 5, 7),
			t:        newIntSet(2, 4, 6, 8),
			expected: `{}`,
		},
	}
	for _, tc := range ts {
		tc.s.IntersectWith(tc.t)
		if got := tc.s.String(); got != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, got)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	ts := []struct {
		s        *IntSet
		t        *IntSet
		expected string
	}{
		{
			s:        newIntSet(1, 3, 5, 7),
			t:        newIntSet(4, 5, 7),
			expected: `{1 3}`,
		},
		{
			s:        newIntSet(1, 3, 5, 7),
			t:        newIntSet(2, 4, 6, 8),
			expected: `{1 3 5 7}`,
		},
	}
	for _, tc := range ts {
		tc.s.DiffereceWith(tc.t)
		if got := tc.s.String(); got != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, got)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	ts := []struct {
		s        *IntSet
		t        *IntSet
		expected string
	}{
		{
			s:        newIntSet(1, 3, 5, 7),
			t:        newIntSet(4, 5, 7),
			expected: `{1 3 4}`,
		},
		{
			s:        newIntSet(1),
			t:        newIntSet(2, 4, 6, 8),
			expected: `{1 2 4 6 8}`,
		},
	}
	for _, tc := range ts {
		tc.s.SymmetricDifference(tc.t)
		if got := tc.s.String(); got != tc.expected {
			t.Errorf("unexpected Len: expected %v, but got %v\n", tc.expected, got)
		}
	}
}
