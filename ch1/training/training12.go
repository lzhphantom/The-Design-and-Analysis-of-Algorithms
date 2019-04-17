package training

func WithLockDoor(n int) (int, int) {
	doors := make([]int, n)
	for i := 1; i <= n; i++ {
		for index := range doors {
			if (index+1)%i == 0 {
				if doors[index] == 0 {
					doors[index] = 1
				} else {
					doors[index] = 0
				}
			}
		}
	}
	var closedDoor int
	for _, value := range doors {
		if value == 0 {
			closedDoor++
		}
	}
	return closedDoor, n - closedDoor
}
