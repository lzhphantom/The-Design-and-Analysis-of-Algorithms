package training

import (
	"math"
	"sort"
)

//It's a bad Algorithms
func MinDistance(arr []int) int {
	dmin := math.MaxInt64
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i]-arr[j] < dmin {
				dmin = arr[i] - arr[j]
			}
		}
	}
	return dmin
}

func MinDistanceUp(arr []int) int {
	dmin := math.MaxInt64
	if !sort.IntsAreSorted(arr) {
		sort.Ints(arr)
	}
	for i := 1; i < len(arr); i++ {
		if dmin == 0 {
			break
		}
		if arr[i]-arr[i-1] < dmin {
			dmin = arr[i] - arr[i-1]
		}
	}
	return dmin
}
