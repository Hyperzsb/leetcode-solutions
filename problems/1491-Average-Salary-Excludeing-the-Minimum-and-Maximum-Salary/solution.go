func average(salary []int) float64 {
	max, min, sum := 0, 10000000, 0
	for _, val := range salary {
		if val > max {
			max = val
		}

		if val < min {
			min = val
		}

		sum += val
	}

	sum -= max + min

	return float64(sum) / float64(len(salary)-2)
}

