func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func minimumAverageDifference(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	prefixSum, minDiff, result := 0, 10000000000, 01
	for i := 0; i < len(nums)-1; i++ {
		prefixSum += nums[i]
		if minDiff > diff(prefixSum/(i+1), (sum-prefixSum)/(len(nums)-i-1)) {
			minDiff = diff(prefixSum/(i+1), (sum-prefixSum)/(len(nums)-i-1))
			result = i
		}
	}
	if minDiff > diff(sum/len(nums), 0) {
		minDiff = diff(sum/len(nums), 0)
		result = len(nums) - 1
	}

	return result
}

