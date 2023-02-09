func totalFruit(fruits []int) int {
	selected := make(map[int]int)

	start, end, max := 0, 0, 0
	for ; end < len(fruits); end++ {
		if selected[fruits[end]] > 0 {
			selected[fruits[end]]++
			continue
		}

		if len(selected) < 2 {
			selected[fruits[end]] = 1
			continue
		}

		if end-start > max {
			max = end - start
		}

		newStart := start
		for i := start; i < end; i++ {
			selected[fruits[i]]--

			if selected[fruits[i]] == 0 {
				newStart = i + 1
				break
			}
		}

		for key, val := range selected {
			if val == 0 {
				delete(selected, key)
			}
		}

		start = newStart

		selected[fruits[end]] = 1
	}

	if end-start > max {
		max = end - start
	}

	return max
}

