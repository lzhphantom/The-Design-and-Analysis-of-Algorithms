package training

func StrIsExist(str string, check string) bool {
	if len(str) < len(check) {
		return false
	}
	indexs := []int{}
	for index, char := range str {
		if char == rune(check[0]) {
			indexs = append(indexs, index)
		}
	}
	for _, index := range indexs {
		if len(str)+1-index >= len(check) {
			if str[index:index+len(check)] == check {
				return true
			}
		} else {
			break
		}
	}
	return false
}
