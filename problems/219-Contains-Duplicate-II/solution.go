func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums)-1 < k {
		k = len(nums) - 1
	}

	set := make(map[int]int)
	idxs, idxe := 0, 0
	for idxe < len(nums) {
		if idxe-idxs <= k {
			if _, exist := set[nums[idxe]]; exist {
				return true
			} else {
				set[nums[idxe]] = 1
			}
			idxe++
		} else {
			delete(set, nums[idxs])
			idxs++
			if _, exist := set[nums[idxe]]; exist {
				return true
			} else {
				set[nums[idxe]] = 1
			}
			idxe++
		}
	}

	return false
}

