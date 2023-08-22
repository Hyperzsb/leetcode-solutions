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

func updateMatrix(mat [][]int) [][]int {
    q := make(Queue, 0)

    for i := range mat {
        for j := range mat[i] {
            if mat[i][j] == 0 {
                q.Push(Point{i, j})
            } else {
                mat[i][j] = math.MaxInt32
            }
        }
    }

    for len(q) > 0 {
        p := q.Front()
        q.Pop()

        if p.R > 0 && mat[p.R-1][p.C] > mat[p.R][p.C] + 1 {
            mat[p.R-1][p.C] = mat[p.R][p.C] + 1
            q.Push(Point{p.R-1, p.C})
        }

        if p.R < len(mat)-1 && mat[p.R+1][p.C] > mat[p.R][p.C] + 1 {
            mat[p.R+1][p.C] = mat[p.R][p.C] + 1
            q.Push(Point{p.R+1, p.C})
        }

        if p.C > 0 && mat[p.R][p.C-1] > mat[p.R][p.C] + 1 {
            mat[p.R][p.C-1] = mat[p.R][p.C] + 1
            q.Push(Point{p.R, p.C-1})
        }

        if p.C < len(mat[0])-1 && mat[p.R][p.C+1] > mat[p.R][p.C] + 1 {
            mat[p.R][p.C+1] = mat[p.R][p.C] + 1
            q.Push(Point{p.R, p.C+1})
        }
    }

    return mat
}

