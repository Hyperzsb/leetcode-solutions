func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	r, c, d, curr := 0, 0, 1, 1
	for {
		result[r][c] = curr
		curr++

		if d == 1 {
			if c < n-1 && result[r][c+1] == 0 {
				c++
			} else if r < n-1 && result[r+1][c] == 0 {
				r++
				d = 2
			} else {
				break
			}
		} else if d == 2 {
			if r < n-1 && result[r+1][c] == 0 {
				r++
			} else if c > 0 && result[r][c-1] == 0 {
				c--
				d = 3
			} else {
				break
			}
		} else if d == 3 {
			if c > 0 && result[r][c-1] == 0 {
				c--
			} else if r > 0 && result[r-1][c] == 0 {
				r--
				d = 4
			} else {
				break
			}
		} else {
			if r > 0 && result[r-1][c] == 0 {
				r--
			} else if c < n-1 && result[r][c+1] == 0 {
				c++
				d = 1
			} else {
				break
			}
		}
	}

	return result
}

