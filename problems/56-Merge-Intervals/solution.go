func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a][0] < intervals[b][0] {
			return true
		} else if intervals[a][0] == intervals[b][0] {
			return intervals[a][1] < intervals[b][1]
		} else {
			return false
		}
	})

	result := make([][]int, 0, len(intervals))
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			result = append(result, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			end = max(end, intervals[i][1])
		}
	}
	result = append(result, []int{start, end})

	return result
}

