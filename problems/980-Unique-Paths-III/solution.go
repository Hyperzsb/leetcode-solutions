func traverse(x, y int, grid *[][]int, path, obstacle *int, result *int) {
	if (*grid)[x][y] == 2 {
		if (*path)+(*obstacle) == len(*grid)*len((*grid)[0])-1 {
			*result += 1
		}
		return
	}

	(*grid)[x][y] = -1
	*path += 1

	if x > 0 && (*grid)[x-1][y] != -1 {
		traverse(x-1, y, grid, path, obstacle, result)
	}

	if x < len(*grid)-1 && (*grid)[x+1][y] != -1 {
		traverse(x+1, y, grid, path, obstacle, result)
	}

	if y > 0 && (*grid)[x][y-1] != -1 {
		traverse(x, y-1, grid, path, obstacle, result)
	}

	if y < len((*grid)[0])-1 && (*grid)[x][y+1] != -1 {
		traverse(x, y+1, grid, path, obstacle, result)
	}

	(*grid)[x][y] = 0
	*path -= 1
}

func uniquePathsIII(grid [][]int) int {
	obstacle, path := 0, 0
	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == -1 {
				obstacle++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				traverse(i, j, &grid, &path, &obstacle, &result)
			}
		}
	}

	return result
}

