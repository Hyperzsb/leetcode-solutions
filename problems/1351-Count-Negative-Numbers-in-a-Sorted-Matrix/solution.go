func countNegatives(grid [][]int) int {
	result := 0
	startCol := 0
	for i := len(grid) - 1; i >= 0; i-- {
		cnt := 0
		for j := len(grid[0]) - 1; j >= startCol; j-- {
			if grid[i][j] >= 0 {
				startCol = j + 1
				break
			}

			cnt++
		}

		result += cnt
	}

	return result
}

