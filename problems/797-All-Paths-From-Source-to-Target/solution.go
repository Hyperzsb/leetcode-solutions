func traverse(idx int, graph *[][]int, current *[]int, result *[][]int) {
	*current = append(*current, idx)
	if idx == len(*graph)-1 {
		newResult := make([]int, len(*current))
		copy(newResult, *current)
		*result = append(*result, newResult)
	} else {
		for i := 0; i < len((*graph)[idx]); i++ {
			traverse((*graph)[idx][i], graph, current, result)
		}
	}
	*current = (*current)[:len(*current)-1]
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	current := make([]int, 0, len(graph))

	traverse(0, &graph, &current, &result)

	return result
}

