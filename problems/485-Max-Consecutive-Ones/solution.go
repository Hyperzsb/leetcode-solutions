func findMaxConsecutiveOnes(nums []int) int {
	max, current := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current++
		}
	}

	if current > max {
		max = current
	}

	return max
}

