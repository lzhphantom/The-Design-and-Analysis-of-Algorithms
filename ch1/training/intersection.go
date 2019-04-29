package training

//两条直线是否有交点

func Intersection(x1, x2, x3, x4, y1, y2, y3, y4 float64) bool {
	var k1, k2 float64
	if x1 != x2 {
		k1 = (y2 - y1) / (x2 - x1)
	}

	if y3 != y4 {
		k2 = (y4 - y3) / (x4 - x3)
	}

	if k1 != k2 {
		return true
	}
	return false
}
