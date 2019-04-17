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
	var t int
	if m > n {
		t = n
	} else {
		t = m
	}
	for t > 0 {
		if m%t == 0 && n%t == 0 {
			break
		}
		t--
	}
	return t, nil
}
