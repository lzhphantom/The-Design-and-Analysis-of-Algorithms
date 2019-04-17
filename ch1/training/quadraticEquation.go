package training

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getAnswer(a, b, c float64) (x1, x2 float64, err error) {
	dt := b*b - 4*a*c
	if dt < 0 {
		return x1, x2, fmt.Errorf("no answer")
	}

	return (-b + math.Sqrt(dt)) / 2 * a, (-b - math.Sqrt(dt)) / 2 * a, nil
}

func DecimalToBinary(n int) string {

	var arr []string
	for n != 1 {
		arr = append(arr, strconv.Itoa(n%2))
		n = n / 2
	}
	arr = append(arr, "1")
	return strings.Join(arr, "")
}
