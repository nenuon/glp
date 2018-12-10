package main

import (
	"bytes"
	"testing"
)

func TestEcho(t *testing.T) {
	var b bytes.Buffer
	Echo(&b, []string{"./echo", "a", "b", "c"})

	got := b.String()
	expected := "./echo a b c\n"
	if got != expected {
		t.Errorf("unexpected output. expected: %v, but got: %v", expected, got)
	}
}
