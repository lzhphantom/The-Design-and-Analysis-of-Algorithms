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
			t.Errorf("SqrtReappear(%d) => %f, want %f", test.m, got, test.want)
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
			t.Errorf("RepeatingNumber(%v) => %v, want %v", test.arr, got, test.want)
		}
	}
}

func TestEuclidExtend(t *testing.T) {
	var tests = []struct {
		m     int
		n     int
		wantx int
		wanty int
	}{
		{60, 24, 1, -2},
		{80, 24, 1, -3},
		{70, 44, 17, -27},
		{50, 34, 15, 22},
	}

	for _, test := range tests {
		gotx, goty := EuclidExtend(test.m, test.n)

		if gotx != test.wantx || goty != test.wanty {
			t.Errorf("EuclidExtend(%d,%d) => %d,%d, want %d,%d",
				test.m, test.n, gotx, goty, test.wantx, test.wanty)
		}

	}
}

func TestWithLockDoor(t *testing.T) {
	var tests = []struct {
		n          int
		wantClosed int
		wantOpened int
	}{
		{5, 3, 2},
		{123, 112, 11},
		{234, 219, 15},
		{2587, 2537, 50},
		{520, 498, 22},
	}

	for _, test := range tests {
		gotClosed, gotOpened := WithLockDoor(test.n)
		if gotClosed != test.wantClosed || gotOpened != test.wantOpened {
			t.Errorf("WithLockDoor(%d) => %d,%d,want %d,%d",
				test.n, gotClosed, gotOpened, test.wantClosed, test.wantOpened)
		}
	}
}
