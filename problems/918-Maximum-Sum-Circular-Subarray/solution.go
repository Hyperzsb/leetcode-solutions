func maxSubarraySumCircular(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	max, currentMax := -1000000000, 0
	min, currentMin := 1000000000, 0
	for i := 0; i < len(nums); i++ {
		currentMax += nums[i]
		currentMin += nums[i]

		if currentMax > max {
			max = currentMax
		}

		if currentMax < 0 {
			currentMax = 0
		}

		if currentMin < min {
			min = currentMin
		}

		if currentMin > 0 {
			currentMin = 0
		}
	}

	if max < 0 {
		return max
	}

	if max > sum-min {
		return max
	} else {
		return sum - min
	}
}

