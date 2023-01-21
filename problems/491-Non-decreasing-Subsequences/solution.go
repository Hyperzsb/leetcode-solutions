func toString(nums *[]int) string {
	str := "#"
	for i := 0; i < len(*nums); i++ {
		str = fmt.Sprintf("%s %d", str, (*nums)[i])
	}

	return str
}

func traverse(current *[]int, start int, nums *[]int, result *[][]int, resultMap *map[string]bool) {
	for i := start; i < len(*nums); i++ {
		if (*nums)[i] >= (*current)[len(*current)-1] {
			*current = append(*current, (*nums)[i])

			str := toString(current)
			if !(*resultMap)[str] {
				(*resultMap)[str] = true

				newResult := make([]int, len(*current))
				copy(newResult, *current)
				*result = append(*result, newResult)

				traverse(current, i+1, nums, result, resultMap)
			}

			*current = (*current)[:len(*current)-1]
		}
	}
}

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	resultMap := make(map[string]bool)
	for i := 0; i < len(nums)-1; i++ {
		current := make([]int, 1)
		current[0] = nums[i]
		traverse(&current, i+1, &nums, &result, &resultMap)
	}

	return result
}

