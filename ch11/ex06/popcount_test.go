package popcount

import (
	"math"
	"math/rand"
	"testing"
)

var ts = []struct {
	x    uint64
	want int
}{
	{x: 1, want: 1},
	{x: 2, want: 1},
	{x: 3, want: 2},
	{x: 1023, want: 10},
}

func TestPopcount(t *testing.T) {
	for _, tc := range ts {
		if got := PopCount(tc.x); got != tc.want {
			t.Errorf("unexpected result: expected %v, but got %v\n", tc.want, got)
		}
	}
}

func TestPopcountLoop(t *testing.T) {
	for _, tc := range ts {
		if got := PopCountLoop(tc.x); got != tc.want {
			t.Errorf("unexpected result: expected %v, but got %v\n", tc.want, got)
		}
	}
}

func TestPopcountBit(t *testing.T) {
	for _, tc := range ts {
		if got := PopCountBit(tc.x); got != tc.want {
			t.Errorf("unexpected result: expected %v, but got %v\n", tc.want, got)
		}
	}
}

func TestPopcountLeastBit(t *testing.T) {
	for _, tc := range ts {
		if got := PopCountLeastBit(tc.x); got != tc.want {
			t.Errorf("unexpected result: expected %v, but got %v\n", tc.want, got)
		}
	}
}

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := uint64(rand.Int63n(math.MaxInt64))
		_ = PopCount(x)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := uint64(rand.Int63n(math.MaxInt64))
		_ = PopCountLoop(x)
	}
}

func BenchmarkPopcountBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := uint64(rand.Int63n(math.MaxInt64))
		_ = PopCountBit(x)
	}
}

func BenchmarkPopcountLeastBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := uint64(rand.Int63n(math.MaxInt64))
		_ = PopCountLeastBit(x)
	}
}
