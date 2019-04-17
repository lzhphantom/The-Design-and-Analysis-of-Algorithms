package training

import (
	"reflect"
	"testing"
)

func TestSqrtReappear(t *testing.T) {
	var tests = []struct {
		m    int
		want float64
	}{
		{25, 5},
		{16, 4},
		{65, 8},
	}

	for _, test := range tests {
		got := SqrtReappear(test.m)
		if got != test.want {
			t.Fatalf("SqrtReappear(%d) => %f, want %f", test.m, got, test.want)
		}
	}
}

func TestRepeatingNumber(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{2, 2, 2, 3, 3, 5, 7, 8}, []int{2, 3}},
		{[]int{1, 2, 2, 3, 3, 5, 7, 7}, []int{2, 3, 7}},
		{[]int{1, 2, 12, 23, 33, 55, 77, 88}, nil},
	}

	for _, test := range tests {
		got := RepeatingNumber(test.arr)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("RepeatingNumber(%v) => %v, want %v", test.arr, got, test.want)
		}
	}
}
