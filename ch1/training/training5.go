package training

//your array must be sorted.
func RepeatingNumber(arr []int) (result []int) {
	var beforeNum int
	for index, num := range arr {
		if index == 0 {
			beforeNum = num
			continue
		}
		if num == beforeNum && !isContain(result, num) {
			result = append(result, num)
		}
		beforeNum = num
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
