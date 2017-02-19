package main

import "testing"

func TestEncore(t *testing.T) {
	slices := make([]*Slice, 3)
	slices[0] = &Slice{X1: 0, Y1: 0, X2: 2, Y2: 1}
	slices[1] = &Slice{X1: 0, Y1: 2, X2: 2, Y2: 2}
	slices[2] = &Slice{X1: 0, Y1: 3, X2: 2, Y2: 4}
	expected := `3
0 0 2 1
0 2 2 2
0 3 2 4`
	output := encode(slices)
	if output != expected {
		t.Errorf("Expected: \n%v\ngot: \n%v", expected, output)
	}
}
