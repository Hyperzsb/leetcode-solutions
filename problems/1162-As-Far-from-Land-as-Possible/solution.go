type Point struct {
	X int
	Y int
}

type Queue []Point

func (q *Queue) Push(x, y int) {
	*q = append(*q, Point{X: x, Y: y})
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() (int, int) {
	return (*q)[0].X, (*q)[0].Y
}

func maxDistance(grid [][]int) int {
	queue := make(Queue, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				queue.Push(i, j)
			}
		}
	}

	result := -1
	for len(queue) > 0 {
		result++

		length := len(queue)
		for i := 0; i < length; i++ {
			x, y := queue.Front()
			queue.Pop()

			if x > 0 && grid[x-1][y] == 0 {
				queue.Push(x-1, y)
				grid[x-1][y] = 1
			}

			if x < len(grid)-1 && grid[x+1][y] == 0 {
				queue.Push(x+1, y)
				grid[x+1][y] = 1
			}

			if y > 0 && grid[x][y-1] == 0 {
				queue.Push(x, y-1)
				grid[x][y-1] = 1
			}

			if y < len(grid[0])-1 && grid[x][y+1] == 0 {
				queue.Push(x, y+1)
				grid[x][y+1] = 1
			}
		}
	}

	if result <= 0 {
		return -1
	} else {
		return result
	}
}

