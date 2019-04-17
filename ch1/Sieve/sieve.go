package Sieve

import "math"

func Sieve(n int) (result []int) {
	var start []int
	for i := 2; i <= n; i++ {
		start = append(start, i)
	}
	p := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= p; i++ {
		if start[i-2] != 0 {
			j := start[i-2] * start[i-2]
			for j <= n {
				start[j-2] = 0
				j += i
			}
		}
	}

	for i := 2; i <= n; i++ {
		if start[i-2] != 0 {
			result = append(result, start[i-2])
		}
	}
	return

}
