package main

import (
	"testing"
)

func TestTupleIsAPoint(t *testing.T) {
	got := tuple{4.3, -4.2, 3.1, 1.0}
	if got.w != 1 {
		t.Errorf("Tuple.w = %f; want 1", got.w)
	}
}

func TestTupleIsNotAPoint(t *testing.T) {
	got := tuple{4.3, -4.2, 3.1, 0.0}
	if got.w != 0 {
		t.Errorf("Tuple.w = %f; want 0", got.w)
	}
}
