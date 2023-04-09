func traverseIsland(r, c int, grid [][]int) bool {
	grid[r][c] = 1

	result := true
	if r > 0 {
		if grid[r-1][c] == 0 {
			re := traverseIsland(r-1, c, grid)
			result = result && re
		}
	} else {
		result = false
	}

	if r < len(grid)-1 {
		if grid[r+1][c] == 0 {
			re := traverseIsland(r+1, c, grid)
			result = result && re
		}
	} else {
		result = false
	}

	if c > 0 {
		if grid[r][c-1] == 0 {
			re := traverseIsland(r, c-1, grid)
			result = result && re
		}
	} else {
		result = false
	}

	if c < len(grid[0])-1 {
		if grid[r][c+1] == 0 {
			re := traverseIsland(r, c+1, grid)
			result = result && re
		}
	} else {
		result = false
	}

	return result
}

func closedIsland(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				if isIsland := traverseIsland(i, j, grid); isIsland {
					result++
				}
			}
		}
	}

	return result
}

