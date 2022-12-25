func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		idx, sum := 0, 0
		for sum < queries[i] && idx < len(nums) {
			sum += nums[idx]
			if sum <= queries[i] {
				idx++
			}
		}
		result[i] = idx
	}

	return result
}

