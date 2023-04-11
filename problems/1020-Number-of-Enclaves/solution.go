func traverse(r, c int, grid [][]int) {
	grid[r][c] = 0

	if r > 0 && grid[r-1][c] == 1 {
		traverse(r-1, c, grid)
	}

	if r < len(grid)-1 && grid[r+1][c] == 1 {
		traverse(r+1, c, grid)
	}

	if c > 0 && grid[r][c-1] == 1 {
		traverse(r, c-1, grid)
	}

	if c < len(grid[0])-1 && grid[r][c+1] == 1 {
		traverse(r, c+1, grid)
	}
}

func numEnclaves(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		if grid[i][0] == 1 {
			traverse(i, 0, grid)
		}

		if grid[i][len(grid[0])-1] == 1 {
			traverse(i, len(grid[0])-1, grid)
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 1 {
			traverse(0, i, grid)
		}

		if grid[len(grid)-1][i] == 1 {
			traverse(len(grid)-1, i, grid)
		}
	}

	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result++
			}
		}
	}

	return result
}

