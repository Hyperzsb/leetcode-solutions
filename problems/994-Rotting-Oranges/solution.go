type Loc struct {
	row int
	col int
}

func orangesRotting(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	empty, fresh, rotten := 0, 0, 0
	rotten_q := make([]Loc, 0)

	for ridx, col := range grid {
		for cidx, cell := range col {
			switch cell {
			case 0:
				empty++
			case 1:
				fresh++
			case 2:
				rotten++
				rotten_q = append(rotten_q, Loc{ridx, cidx})
			}
		}
	}

	timer := 0
	for len(rotten_q) > 0 {
		if fresh == 0 {
			break
		}

		length := len(rotten_q)

		for i := 0; i < length; i++ {
			front := rotten_q[i]

			if front.row > 0 {
				if grid[front.row-1][front.col] == 1 {
					fresh--
					grid[front.row-1][front.col] = 2
					rotten_q = append(rotten_q, Loc{front.row - 1, front.col})
				}
			}

			if front.row < row-1 {
				if grid[front.row+1][front.col] == 1 {
					fresh--
					grid[front.row+1][front.col] = 2
					rotten_q = append(rotten_q, Loc{front.row + 1, front.col})
				}
			}

			if front.col > 0 {
				if grid[front.row][front.col-1] == 1 {
					fresh--
					grid[front.row][front.col-1] = 2
					rotten_q = append(rotten_q, Loc{front.row, front.col - 1})
				}
			}

			if front.col < col-1 {
				if grid[front.row][front.col+1] == 1 {
					fresh--
					grid[front.row][front.col+1] = 2
					rotten_q = append(rotten_q, Loc{front.row, front.col + 1})
				}
			}
		}

		rotten_q = rotten_q[length:]

		timer++
	}

	if fresh > 0 {
		return -1
	} else {
		return timer
	}
}

