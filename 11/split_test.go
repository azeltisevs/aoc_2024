package main

import "testing"

func TestSplitNumber(t *testing.T) {
	l, r := splitNumber(17, 2)

	if l != 123 || r != 456 {
		t.Error("wrong output", l, r)
	}
}
