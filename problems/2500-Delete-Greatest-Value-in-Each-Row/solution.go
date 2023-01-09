func deleteGreatestValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		sort.Slice(grid[i], func(a, b int) bool {
			return grid[i][a] > grid[i][b]
		})
	}

	result := 0
	for i := 0; i < len(grid[0]); i++ {
		max := 0
		for j := 0; j < len(grid); j++ {
			if max < grid[j][i] {
				max = grid[j][i]
			}
		}
		result += max
	}

	return result
}

