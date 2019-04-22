package training

//比较计数排序
func ComparisonCountingSort(arr []int) []int {
	var outputSort = make([]int, len(arr))
	Count := make([]int, len(arr))

	for i := 0; i <= len(arr)-2; i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] < arr[j] {
				Count[j] = Count[j] + 1
			} else {
				Count[i] = Count[i] + 1
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		outputSort[Count[i]] = arr[i]
	}

	return outputSort
}
