func binarySearch(nums []int, target int) int {
	start, half, end := 0, (len(nums)-1)/2, len(nums)-1
	for start < end {
		if nums[half] <= target {
			start = half + 1
		} else {
			end = half
		}

		half = (start + end) / 2
	}

	return start
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make([]int, len(obstacles))
	tail := make([]int, 1)

	tail[0] = obstacles[0]
	result[0] = 1

	for i := 1; i < len(obstacles); i++ {
		if obstacles[i] >= tail[len(tail)-1] {
			tail = append(tail, obstacles[i])
			result[i] = len(tail)
		} else {
			idx := binarySearch(tail, obstacles[i])
			tail[idx] = obstacles[i]
			result[i] = idx + 1
		}
	}

	return result
}

