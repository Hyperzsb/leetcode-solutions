type Point struct {
    R int
    C int
}

type Queue []Point

func (q *Queue) Push(p Point) {
    *q = append(*q, p)
}

func (q *Queue) Pop() {
    *q = (*q)[1:]
}

func (q *Queue) Front() Point {
    return (*q)[0]
}

func isReachable(row, col int, cells [][]int, day int) bool {
    grid := make([][]bool, row)
    for i := range grid {
        grid[i] = make([]bool, col)
    }

    for i := 0; i < day; i++ {
        grid[cells[i][0]-1][cells[i][1]-1] = true
    }

    q := make(Queue, 0)
    for i := 0; i < col; i++ {
        if !grid[0][i] {
            q.Push(Point{0, i})
        }
    }

    for len(q) > 0 {
        length := len(q)
        for i := 0; i < length; i++ {
            p := q.Front()
            q.Pop()

            if p.R < 0 || p.R > row-1 || p.C < 0 || p.C > col-1 {
                continue
            }

            if grid[p.R][p.C] {
                continue
            }

            if p.R == row-1 {
                return true
            }

            grid[p.R][p.C] = true

            q.Push(Point{p.R-1, p.C})
            q.Push(Point{p.R+1, p.C})
            q.Push(Point{p.R, p.C-1})
            q.Push(Point{p.R, p.C+1})
        }
    }

    return false
}

func latestDayToCross(row int, col int, cells [][]int) int {
    start, end := 0, row * col
    for start < end {
        half := (start + end) / 2

        if isReachable(row, col, cells, half) {
            start = half + 1
        } else {
            end = half
        }
    }

    return start-1
}

