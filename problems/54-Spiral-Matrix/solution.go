func spiralOrder(matrix [][]int) []int {
	row, r, col, c, direction := len(matrix), 0, len(matrix[0]), 0, 1
	dummy := 1000

	result := make([]int, 0, row*col)
	for {
		result = append(result, matrix[r][c])
		matrix[r][c] = dummy

		if direction == 1 {
			if c < col-1 && matrix[r][c+1] != dummy {
				c++
			} else if r < row-1 && matrix[r+1][c] != dummy {
				r++
				direction = 2
			} else {
				break
			}
		} else if direction == 2 {
			if r < row-1 && matrix[r+1][c] != dummy {
				r++
			} else if c > 0 && matrix[r][c-1] != dummy {
				c--
				direction = 3
			} else {
				break
			}
		} else if direction == 3 {
			if c > 0 && matrix[r][c-1] != dummy {
				c--
			} else if r > 0 && matrix[r-1][c] != dummy {
				r--
				direction = 4
			} else {
				break
			}
		} else {
			if r > 0 && matrix[r-1][c] != dummy {
				r--
			} else if c < col-1 && matrix[r][c+1] != dummy {
				c++
				direction = 1
			} else {
				break
			}
		}
	}

	return result
}

