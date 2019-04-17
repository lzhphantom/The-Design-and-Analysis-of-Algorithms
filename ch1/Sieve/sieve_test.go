package Sieve

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	var tests = []struct {
		n    int
		want []int
	}{
		{25, []int{2, 3, 5, 7, 11, 13, 17, 19, 23}},
		{3, []int{2, 3}},
		{9, []int{2, 3, 5, 7, 8}},
	}

	for _, test := range tests {
		got := Sieve(test.n)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Sieve(%d) => %v, want %v", test.n, got, test.want)
		}
	}
}
