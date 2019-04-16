package Euclid

func Euclid(m, n int) int {
	for n != 0 {
		r := m % n
		m = n
		n = r
	}
	return m
}
