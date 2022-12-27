func search(nums []int, target int) int {
	start, half, end := 0, (len(nums)-1)/2, len(nums)-1
	for start < end {
		if target <= nums[half] {
			end = half
		} else {
			start = half + 1
		}
		half = (start + end) / 2
	}

	if target == nums[half] {
		return half
	} else {
		return -1
	}
}

