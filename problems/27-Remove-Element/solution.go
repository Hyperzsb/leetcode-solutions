func removeElement(nums []int, val int) int {
	idxa, idxb := 0, 0
	for idxb < len(nums) {
		if nums[idxb] != val {
			nums[idxa] = nums[idxb]
			idxa++
			idxb++
		} else {
			idxb++
		}
	}

	return idxa
}

