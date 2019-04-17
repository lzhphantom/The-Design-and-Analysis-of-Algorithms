package training

import (
	"fmt"
	"math"
)

type Triangle struct {
	a float64
	b float64
	c float64
}

func NewTriangle(a, b, c float64) (*Triangle, error) {
	if c <= a+b {
		return nil, fmt.Errorf("against the rule of triangle:%f+%f<=%f", a, b, c)
	}
	return &Triangle{a, b, c}, nil
}

func (t *Triangle) Area() float64 {
	p := (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}

func (t *Triangle) AreaWithTrifunc() float64 {
	sinA := math.Sqrt(1 - math.Pow((t.b*t.b+t.c*t.c-t.a*t.a)/2*t.b*t.c, 2))
	return t.b * t.c * sinA / 2
}
