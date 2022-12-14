func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) >= 3 {
		nums[2] += nums[0]
	}

	for i := 3; i < len(nums); i++ {
		if nums[i-2] >= nums[i-3] {
			nums[i] += nums[i-2]
		} else {
			nums[i] += nums[i-3]
		}
	}

	if nums[len(nums)-1] >= nums[len(nums)-2] {
		return nums[len(nums)-1]
	} else {
		return nums[len(nums)-2]
	}
}

