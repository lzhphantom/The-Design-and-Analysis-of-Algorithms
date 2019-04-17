package training

//your array must be sorted.
func RepeatingNumber(arr []int) (result []int) {
	var bfNum int
	for index, num := range arr {
		if index == 0 {
			bfNum = num
			continue
		}
		if num == bfNum && !isContain(result, num) {
			result = append(result, num)
		}
		bfNum = num
	}
	return
}

func isContain(arr []int, checkNum int) bool {
	for _, num := range arr {
		if num == checkNum {
			return true
		}
	}
	return false
}
