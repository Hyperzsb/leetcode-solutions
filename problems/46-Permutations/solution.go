func traverse(nums []int, current *[]int, result *[][]int) {
	if len(nums) == 1 {
		*current = append(*current, nums[0])
		newResult := make([]int, len(*current))
		copy(newResult, *current)
		*result = append(*result, newResult)
		*current = (*current)[:len(*current)-1]
		return
	}
	for i := 0; i < len(nums); i++ {
		*current = append(*current, nums[i])

		newNums := make([]int, i)
		copy(newNums, nums[:i])
		newNums = append(newNums, nums[i+1:]...)

		traverse(newNums, current, result)
		*current = (*current)[:len(*current)-1]
	}
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0, len(nums))

	traverse(nums, &current, &result)

	return result
}

