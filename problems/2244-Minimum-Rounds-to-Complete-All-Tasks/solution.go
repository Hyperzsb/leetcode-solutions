func minimumRounds(tasks []int) int {
	taskMap := make(map[int]int)
	for _, val := range tasks {
		taskMap[val]++
	}

	result := 0
	for _, val := range taskMap {
		if val == 1 {
			return -1
		}
		if val%3 != 0 {
			result += val/3 + 1
		} else {
			result += val / 3
		}
	}

	return result
}

