func findMax(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func minimizeArrayValue(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	surfixSum := 0
	for i := len(nums) - 1; i > 0; i-- {
		average := 0
		if (sum-surfixSum)%(i+1) == 0 {
			average = (sum - surfixSum) / (i + 1)
		} else {
			average = (sum-surfixSum)/(i+1) + 1
		}

		if nums[i] > average {
			nums[i-1] += nums[i] - average
			nums[i] -= nums[i] - average
		}

		surfixSum += nums[i]
	}

	return findMax(nums)
}

