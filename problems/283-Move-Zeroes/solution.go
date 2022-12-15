func moveZeroes(nums []int) {
	zeroIdx, nonzeroIdx := 0, 0

	for {
		for zeroIdx < len(nums) && nums[zeroIdx] != 0 {
			zeroIdx++
		}

		if zeroIdx == len(nums) {
			break
		}

		nonzeroIdx = zeroIdx + 1

		for nonzeroIdx < len(nums) && nums[nonzeroIdx] == 0 {
			nonzeroIdx++
		}

		if nonzeroIdx == len(nums) {
			break
		}

		nums[zeroIdx] = nums[nonzeroIdx]
		nums[nonzeroIdx] = 0
	}
}

