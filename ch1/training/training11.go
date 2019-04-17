package training

import "github.com/lzhphantom/The-Design-and-Analysis-of-Algorithms/ch1/Euclid"

//get one answer EuclidExtend
func EuclidExtend(m, n int) (x, y int) {
	commonDivisor, _ := Euclid.Euclid(m, n)
	for true {
		if (commonDivisor-m*x)%n == 0 {
			y = (commonDivisor - m*x) / n
			break
		}
		x++
	}
	return
}
