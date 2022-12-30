func traverse(idx int, candidates *[]int, target int, result *[][]int, current *[]int) {
	count := 0
	for (*candidates)[idx]*count <= target {
		if (*candidates)[idx]*count == target {
			new := make([]int, len(*current))
			copy(new, *current)
			*result = append(*result, new)
			break
		}
		if idx < len(*candidates)-1 {
			traverse(idx+1, candidates, target-(*candidates)[idx]*count, result, current)
		}
		count++
		*current = append(*current, (*candidates)[idx])
	}
	(*current) = (*current)[:len(*current)-count]
}

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0)

	traverse(0, &candidates, target, &result, &current)

	return result
}

