func equalPairs(grid [][]int) int {
	rowList := make([]string, len(grid))
	for i := range grid {
		row := ""
		for j := range grid[0] {
			row = fmt.Sprintf("%s%d.", row, grid[i][j])
		}
		rowList[i] = row
	}

	colMap := make(map[string]int)
	for i := range grid[0] {
		col := ""
		for j := range grid {
			col = fmt.Sprintf("%s%d.", col, grid[j][i])
		}
		colMap[col]++
	}

	result := 0
	for i := range grid {
		if cnt, ok := colMap[rowList[i]]; ok {
			result += cnt
		}
	}

	return result
}

