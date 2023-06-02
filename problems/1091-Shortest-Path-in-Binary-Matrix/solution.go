type MNode struct {
	X int
	Y int
}

type Queue []MNode

func (q *Queue) Push(n MNode) {
	*q = append(*q, n)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() MNode {
	return (*q)[0]
}

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	q := make(Queue, 0)
	q.Push(MNode{0, 0})
	grid[0][0] = -1

	result, flag := 0, false
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			cn := q.Front()
			q.Pop()

			if cn.X == len(grid)-1 && cn.Y == len(grid[0])-1 {
				flag = true
				break
			}

			if cn.X > 0 && cn.Y > 0 && grid[cn.X-1][cn.Y-1] == 0 {
				q.Push(MNode{cn.X - 1, cn.Y - 1})
				grid[cn.X-1][cn.Y-1] = -1
			}

			if cn.X > 0 && grid[cn.X-1][cn.Y] == 0 {
				q.Push(MNode{cn.X - 1, cn.Y})
				grid[cn.X-1][cn.Y] = -1
			}

			if cn.X > 0 && cn.Y < len(grid[0])-1 && grid[cn.X-1][cn.Y+1] == 0 {
				q.Push(MNode{cn.X - 1, cn.Y + 1})
				grid[cn.X-1][cn.Y+1] = -1
			}

			if cn.Y < len(grid[0])-1 && grid[cn.X][cn.Y+1] == 0 {
				q.Push(MNode{cn.X, cn.Y + 1})
				grid[cn.X][cn.Y+1] = -1
			}

			if cn.X < len(grid)-1 && cn.Y < len(grid[0])-1 && grid[cn.X+1][cn.Y+1] == 0 {
				q.Push(MNode{cn.X + 1, cn.Y + 1})
				grid[cn.X+1][cn.Y+1] = -1
			}

			if cn.X < len(grid)-1 && grid[cn.X+1][cn.Y] == 0 {
				q.Push(MNode{cn.X + 1, cn.Y})
				grid[cn.X+1][cn.Y] = -1
			}

			if cn.X < len(grid)-1 && cn.Y > 0 && grid[cn.X+1][cn.Y-1] == 0 {
				q.Push(MNode{cn.X + 1, cn.Y - 1})
				grid[cn.X+1][cn.Y-1] = -1
			}

			if cn.Y > 0 && grid[cn.X][cn.Y-1] == 0 {
				q.Push(MNode{cn.X, cn.Y - 1})
				grid[cn.X][cn.Y-1] = -1
			}
		}

		result++

		if flag {
			break
		}
	}

	if flag {
		return result
	} else {
		return -1
	}
}

