func findDisappearedNumbers(nums []int) []int {
	numsAppeared := make([]bool, len(nums)+1)
	numsAppeared[0] = true

	for _, val := range nums {
		numsAppeared[val] = true
	}

	result := make([]int, 0)
	for idx, val := range numsAppeared {
		if !val {
			result = append(result, idx)
		}
	}

	return result
}

