type Point struct {
	X int
	Y int
}

type Queue []Point

func (q *Queue) Push(p Point) {
	*q = append(*q, p)
}

func (q *Queue) Front() Point {
	return (*q)[0]
}

func (q *Queue) Pop() {
	if cap(*q) > 2*len(*q) {
		nq := make(Queue, len(*q)-1)
		copy(nq, (*q)[1:])
		*q = nq
	} else {
		(*q) = (*q)[1:]
	}
}

func traverse(p Point, grid [][]int, coast *Queue) {
	grid[p.X][p.Y] = 2

	flag := true
	if p.X > 0 {
		if grid[p.X-1][p.Y] == 0 && flag {
			coast.Push(p)
			flag = false
		}

		if grid[p.X-1][p.Y] == 1 {
			traverse(Point{p.X - 1, p.Y}, grid, coast)
		}
	}

	if p.X < len(grid)-1 {
		if grid[p.X+1][p.Y] == 0 && flag {
			coast.Push(p)
			flag = false
		}

		if grid[p.X+1][p.Y] == 1 {
			traverse(Point{p.X + 1, p.Y}, grid, coast)
		}
	}

	if p.Y > 0 {
		if grid[p.X][p.Y-1] == 0 && flag {
			coast.Push(p)
			flag = false
		}

		if grid[p.X][p.Y-1] == 1 {
			traverse(Point{p.X, p.Y - 1}, grid, coast)
		}
	}

	if p.Y < len(grid[0])-1 {
		if grid[p.X][p.Y+1] == 0 && flag {
			coast.Push(p)
			flag = false
		}

		if grid[p.X][p.Y+1] == 1 {
			traverse(Point{p.X, p.Y + 1}, grid, coast)
		}
	}
}

func findShortest(grid [][]int, coast *Queue) int {
	result := 0

	for len(*coast) > 0 {
		length := len(*coast)
		for i := 0; i < length; i++ {
			cp := coast.Front()
			coast.Pop()

			if cp.X > 0 {
				if grid[cp.X-1][cp.Y] == 1 {
					return result
				}

				if grid[cp.X-1][cp.Y] == 0 {
					grid[cp.X-1][cp.Y] = 2
					coast.Push(Point{cp.X - 1, cp.Y})
				}
			}

			if cp.X < len(grid)-1 {
				if grid[cp.X+1][cp.Y] == 1 {
					return result
				}

				if grid[cp.X+1][cp.Y] == 0 {
					grid[cp.X+1][cp.Y] = 2
					coast.Push(Point{cp.X + 1, cp.Y})
				}
			}

			if cp.Y > 0 {
				if grid[cp.X][cp.Y-1] == 1 {
					return result
				}

				if grid[cp.X][cp.Y-1] == 0 {
					grid[cp.X][cp.Y-1] = 2
					coast.Push(Point{cp.X, cp.Y - 1})
				}
			}

			if cp.Y < len(grid[0])-1 {
				if grid[cp.X][cp.Y+1] == 1 {
					return result
				}

				if grid[cp.X][cp.Y+1] == 0 {
					grid[cp.X][cp.Y+1] = 2
					coast.Push(Point{cp.X, cp.Y + 1})
				}
			}
		}

		result++
	}

	return result
}

func shortestBridge(grid [][]int) int {
	coast := make(Queue, 0)

	flag := false
	for i := 0; i < len(grid); i++ {
		if flag {
			break
		}

		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				traverse(Point{i, j}, grid, &coast)
				flag = true
				break
			}
		}
	}

	return findShortest(grid, &coast)
}

