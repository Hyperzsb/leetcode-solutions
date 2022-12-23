func traverse(r, c int, grid *[][]int) int {
	if (*grid)[r][c] == -1 {
		return 0
	}
	if (*grid)[r][c] == 0 {
		return 1
	}

	(*grid)[r][c] = -1
	result := 0
	if r <= 0 {
		result++
	} else {
		result += traverse(r-1, c, grid)
	}

	if r >= len(*grid)-1 {
		result++
	} else {
		result += traverse(r+1, c, grid)
	}

	if c <= 0 {
		result++
	} else {
		result += traverse(r, c-1, grid)
	}

	if c >= len((*grid)[0])-1 {
		result++
	} else {
		result += traverse(r, c+1, grid)
	}

	return result
}

func islandPerimeter(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result = traverse(i, j, &grid)
			}
		}
	}

	return result
}

