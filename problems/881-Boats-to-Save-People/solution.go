func search(target int, arr []int) int {
	start, end := 0, len(arr)-1
	for start < end {
		half := (start + end) / 2

		if target <= arr[half] {
			end = half
		} else {
			start = half + 1
		}
	}

	return (start + end) / 2
}

func numRescueBoats(people []int, limit int) int {
	sort.Slice(people, func(a, b int) bool {
		return people[a] < people[b]
	})

	visited := make([]bool, len(people))

	result := 0
	for i := 0; i < len(people); i++ {
		if visited[i] {
			continue
		}

		if 2*people[i] > limit {
			result++
			continue
		}

		visited[i] = true

		findIdx, maxIdx := -1, search(limit-people[i]+1, people)
		for j := maxIdx; j > i; j-- {
			if !visited[j] && people[j]+people[i] <= limit {
				findIdx = j
				break
			}
		}

		if findIdx != -1 {
			visited[findIdx] = true
		}

		result++
	}

	return result
}

