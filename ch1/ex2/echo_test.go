package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	var b bytes.Buffer
	// 入力
	Echo(&b, []string{"./echo", "a", "b", "c"})

	got := b.String()
	expected := "0 a\n1 b\n2 c\n"
	if got != expected {
		t.Errorf("Test failed. expected: %v, got %v", expected, got)
	}
}
