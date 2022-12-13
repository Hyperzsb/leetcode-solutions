func containsDuplicate(nums []int) bool {
	set := make(map[int]int)
	for _, val := range nums {
		if _, exist := set[val]; !exist {
			set[val] = 1
		} else {
			set[val]++
		}
	}

	return len(set) < len(nums)
}

