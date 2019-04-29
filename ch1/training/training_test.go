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

func TestDecimalToBinary(t *testing.T) {
	var tests = []struct {
		dec  int
		want string
	}{
		{10, "1010"},
	}
	for _, test := range tests {
		got := DecimalToBinary(test.dec)
		if got != test.want {
			t.Errorf("DecimalToBinary(%d) => %s, want %s", test.dec, got, test.want)
		}
	}
}

func TestMinDistance(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1},
		{[]int{21, 12, 123, 43, 522, 61, 37, 83, 92}, 6},
		{[]int{221, 12 << 2, 12 << 3, 4 << 3, 52 << 2, 621, 3 << 7, 8 << 3, 9 << 2}, 4},
	}

	for _, test := range tests {
		got := MinDistance(test.arr)
		if got != test.want {
			t.Errorf("MinDistance(%v) => %v,want %v", test.arr, got, test.want)
		}
	}
}

func TestMinDistanceUp(t *testing.T) {
	var tests = []struct {
		arr  []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1},
		{[]int{21, 12, 123, 43, 522, 61, 37, 83, 92}, 6},
		{[]int{221, 12 << 2, 12 << 3, 4 << 3, 52 << 2, 621, 3 << 7, 8 << 3, 9 << 2}, 4},
	}

	for _, test := range tests {
		got := MinDistanceUp(test.arr)
		if got != test.want {
			t.Errorf("MinDistanceUp(%v) => %v,want %v", test.arr, got, test.want)
		}
	}
}

func TestComparisonCountingSort(t *testing.T) {
	var tests = []struct {
		arr  []int
		want []int
	}{
		{[]int{60, 35, 81, 98, 14, 47}, []int{14, 35, 47, 60, 81, 98}},
	}

	for _, test := range tests {
		got := ComparisonCountingSort(test.arr)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("ComparisonCountingSort(%v) => %v,want %v", test.arr, got, test.want)
		}
	}
}

func TestStrIsExist(t *testing.T) {
	var tests = []struct {
		str   string
		check string
		want  bool
	}{
		{"123", "12", true},
		{"3359", "39", false},
		{"12sif", "2f", false},
		{"1", "12", false},
		{"good boy", "od", true},
		{"good boy", "db", false},
	}
	for _, test := range tests {
		got := StrIsExist(test.str, test.check)
		if got != test.want {
			t.Errorf("StrIsExist(%s,%s) => %v,want %v", test.str, test.check, got, test.want)
		}
	}
}

func TestStrIsExistUp(t *testing.T) {
	var tests = []struct {
		str   string
		check string
		want  bool
	}{
		{"123", "12", true},
		{"3359", "39", false},
		{"12sif", "2f", false},
		{"1", "12", false},
		{"good boy", "od", true},
		{"good boy", "db", false},
	}
	for _, test := range tests {
		got := StrIsExistUp(test.str, test.check)
		if got != test.want {
			t.Errorf("StrIsExistUp(%s,%s) => %v,want %v", test.str, test.check, got, test.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	var tests = []struct {
		x1, x2, x3, x4 float64
		y1, y2, y3, y4 float64
		want           bool
	}{
		{1, 1, 1, 1, 1, 1, 1, 1, false},
		{1, 2, 1, 3, 1, 5, 1, 7, true},
		{3, 4, 5, 6, 11, 12, 13, 14, false},
		{13, -21, -1, 12, -1, 20, 10, -21, true},
		{123, 456, 789, 357, 321, 654, 987, 753, true},
		{21, 31, 41, 51, 19, 18, 17, 16, false},
	}

	for _, test := range tests {
		got := Intersection(test.x1, test.x2, test.x3, test.x4, test.y1, test.y2, test.y3, test.y4)

		if got != test.want {
			t.Errorf("Intersection(%f,%f,%f,%f,%f,%f,%f,%f) => %v,want %v",
				test.x1, test.x2, test.x3, test.x4, test.y1, test.y2, test.y3, test.y4, got, test.want)
		}
	}
}
