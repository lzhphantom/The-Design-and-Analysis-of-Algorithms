package training

func SqrtReappear(m int) float64 {
	if m == 0 {
		return 0
	}

	min := float64(1)
	max := float64(m / 2)
	if min*min == float64(m) {
		return min
	}
	if max*max == float64(m) {
		return max
	}

	return sqrtValue(max, min, m)
}

func sqrtValue(max, min float64, m int) float64 {
	n := (max + min) / 2
	value := n * n
	if value < float64(m) {
		n = sqrtValue(max, n, m)
	} else if value > float64(m) {
		n = sqrtValue(n, min, m)
	}
	return n
}
