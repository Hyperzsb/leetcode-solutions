func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := make([][]int, 0, len(intervals)+1)

	if newInterval[1] < intervals[0][0] {
		result = append(result, newInterval)
		result = append(result, intervals...)
		return result
	}
	if newInterval[1] <= intervals[0][1] {
		intervals[0][0] = min(newInterval[0], intervals[0][0])
		result = append(result, intervals...)
		return result
	}

	if newInterval[0] > intervals[len(intervals)-1][1] {
		result = append(result, intervals...)
		result = append(result, newInterval)
		return result
	}
	if newInterval[0] >= intervals[len(intervals)-1][0] {
		intervals[len(intervals)-1][1] =
			max(newInterval[1], intervals[len(intervals)-1][1])
		result = append(result, intervals...)
		return result
	}

	start := -1
	for i := 0; i < len(intervals); i++ {
		if newInterval[0] < intervals[i][0] ||
			(newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i][1]) {
			start = i
			break
		}
	}

	end := -1
	for i := len(intervals) - 1; i >= 0; i-- {
		if newInterval[1] > intervals[i][1] ||
			(newInterval[1] <= intervals[i][1] && newInterval[1] >= intervals[i][0]) {
			end = i
			break
		}
	}

	if start > end {
		result = append(result, intervals[:end+1]...)
		result = append(result, newInterval)
		result = append(result, intervals[start:]...)
		return result
	} else {
		result = append(result, intervals[:start]...)
		newInterval[0] = min(newInterval[0], intervals[start][0])
		newInterval[1] = max(newInterval[1], intervals[end][1])
		result = append(result, newInterval)
		result = append(result, intervals[end+1:]...)
	}

	return result
}

