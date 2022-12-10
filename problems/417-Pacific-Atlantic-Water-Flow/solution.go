type Loc struct {
	r int
	c int
}

func pacificAtlantic(heights [][]int) [][]int {
	row, col := len(heights), len(heights[0])
	pacific_flags := make([][]bool, row)
	atlantic_flags := make([][]bool, row)
	for i := 0; i < row; i++ {
		pacific_flags[i] = make([]bool, col)
		atlantic_flags[i] = make([]bool, col)
		for j := 0; j < col; j++ {
			pacific_flags[i][j] = false
			atlantic_flags[i][j] = false
		}
	}

	pacific_queue := make([]Loc, 0)
	atlantic_queue := make([]Loc, 0)

	for i := 0; i < row; i++ {
		pacific_flags[i][0] = true
		atlantic_flags[i][col-1] = true
		pacific_queue = append(pacific_queue, Loc{i, 0})
		atlantic_queue = append(atlantic_queue, Loc{i, col - 1})
	}

	for i := 0; i < col; i++ {
		pacific_flags[0][i] = true
		atlantic_flags[row-1][i] = true
		pacific_queue = append(pacific_queue, Loc{0, i})
		atlantic_queue = append(atlantic_queue, Loc{row - 1, i})
	}

	for i := 0; i < len(pacific_queue); i++ {
		front := pacific_queue[i]
		r, c := front.r, front.c

		pacific_flags[r][c] = true

		if r > 0 && heights[r-1][c] >= heights[r][c] && !pacific_flags[r-1][c] {
			pacific_queue = append(pacific_queue, Loc{r - 1, c})
		}

		if r < row-1 && heights[r+1][c] >= heights[r][c] && !pacific_flags[r+1][c] {
			pacific_queue = append(pacific_queue, Loc{r + 1, c})
		}

		if c > 0 && heights[r][c-1] >= heights[r][c] && !pacific_flags[r][c-1] {
			pacific_queue = append(pacific_queue, Loc{r, c - 1})
		}

		if c < col-1 && heights[r][c+1] >= heights[r][c] && !pacific_flags[r][c+1] {
			pacific_queue = append(pacific_queue, Loc{r, c + 1})
		}
	}

	for i := 0; i < len(atlantic_queue); i++ {
		front := atlantic_queue[i]
		r, c := front.r, front.c

		atlantic_flags[r][c] = true

		if r > 0 && heights[r-1][c] >= heights[r][c] && !atlantic_flags[r-1][c] {
			atlantic_queue = append(atlantic_queue, Loc{r - 1, c})
		}

		if r < row-1 && heights[r+1][c] >= heights[r][c] && !atlantic_flags[r+1][c] {
			atlantic_queue = append(atlantic_queue, Loc{r + 1, c})
		}

		if c > 0 && heights[r][c-1] >= heights[r][c] && !atlantic_flags[r][c-1] {
			atlantic_queue = append(atlantic_queue, Loc{r, c - 1})
		}

		if c < col-1 && heights[r][c+1] >= heights[r][c] && !atlantic_flags[r][c+1] {
			atlantic_queue = append(atlantic_queue, Loc{r, c + 1})
		}
	}

	result := make([][]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pacific_flags[i][j] && atlantic_flags[i][j] {
				loc := make([]int, 2)
				loc[0] = i
				loc[1] = j
				result = append(result, loc)
			}
		}
	}

	return result
}

