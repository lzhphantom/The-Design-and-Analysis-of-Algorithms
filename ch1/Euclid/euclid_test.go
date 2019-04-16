package Euclid

import (
	"testing"
)

func TestEuclid(t *testing.T) {
	var tests = []struct {
		m    int
		n    int
		want int
	}{
		{60, 24, 12},
		{33, 21, 3},
		{1238, 125, 1},
		{123456789, 456789, 3},
		{987654321, 123456789, 9},
	}
	for _, test := range tests {
		got := Euclid(test.m, test.n)

		if got != test.want {
			t.Errorf("Euclid(%d,%d)=>%d, want %d",
				test.m, test.n, got, test.want)
		}
	}
}
