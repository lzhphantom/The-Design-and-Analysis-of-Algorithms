package Euclid

import "fmt"

func Euclid(m, n int) (int, error) {
	if m == 0 || n == 0 {
		return 0, fmt.Errorf("%d doesn't have the greatest common divisor with any natural number", 0)
	}
	for n != 0 {
		r := m % n
		m = n
		n = r
	}
	return m, nil
}

func GcdSerialInt(m, n int) (int, error) {
	if m == 0 || n == 0 {
		return 0, fmt.Errorf("%d doesn't have the greatest common divisor with any natural number", 0)
	}
	var min int
	if m > n {
		min = n
	} else {
		min = m
	}
	for min > 0 {
		if m%min == 0 && n%min == 0 {
			break
		}
		min--
	}
	return min, nil
}
