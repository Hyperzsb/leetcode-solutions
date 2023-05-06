func numSubseq(nums []int, target int) int {
	powers := make([]int, len(nums))
	powers[0] = 1
	for i := 1; i < len(nums); i++ {
		powers[i] = (powers[i-1] * 2) % (1e9 + 7)
	}

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	result := 0
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			result = (result + powers[right-left]) % (1e9 + 7)
			left++
		}
	}

	return result
}

