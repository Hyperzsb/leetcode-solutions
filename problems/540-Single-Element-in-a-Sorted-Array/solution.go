func singleNonDuplicate(nums []int) int {
	previous := -1

	for i := 0; i < len(nums); i++ {
		if previous == -1 {
			previous = nums[i]
			continue
		}

		if nums[i] == previous {
			previous = -1
		} else {
			break
		}
	}

	return previous
}

