func searchInsert(nums []int, target int) int {
	start, half, end := 0, 0, len(nums)-1

	for start < end {
		half = (start + end) / 2

		if nums[half] < target {
			start = half + 1
		} else if nums[half] == target {
			return half
		} else {
			end = half
		}
	}

	half = (start + end) / 2
	if nums[half] < target {
		return half + 1
	} else if nums[half] == target {
		return half
	} else {
		return half
	}
}

