func missingNumber(nums []int) int {
	targetSum := (1 + len(nums)) * len(nums) / 2
	actualSum := 0
	for _, val := range nums {
		actualSum += val
	}

	return targetSum - actualSum
}

