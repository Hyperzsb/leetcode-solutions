type Point struct {
    R int
    C int
    M int
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

func shortestPathAllKeys(grid []string) int {
    keyBitMap := make(map[byte]int)
    klCnt := 0
    sR, sC := -1, -1
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] >= 'a' && grid[i][j] <= 'z' {
                keyBitMap[grid[i][j]] = klCnt
                keyBitMap[grid[i][j]-'a'+'A'] = klCnt
                klCnt++
                continue
            }

            if grid[i][j] == '@' {
                sR, sC = i, j
            }
        }
    }

    dp := make([][][]bool, len(grid))
    for i := range dp {
        dp[i] = make([][]bool, len(grid[0]))
        for j := range dp[i] {
            dp[i][j] = make([]bool, 1 << klCnt)
        }
    }

    q := make(Queue, 0)
    q.Push(Point{sR, sC, 0})

    result := 0
    for len(q) > 0 {
        length := len(q)
        for i := 0; i < length; i++ {
            p := q.Front()
            q.Pop()

            if p.R < 0 || p.R > len(grid)-1 || p.C < 0 || p.C > len(grid[0])-1 {
                continue
            }

            if grid[p.R][p.C] == '#' {
                continue
            }

            if grid[p.R][p.C] >= 'A' && grid[p.R][p.C] <= 'Z' {
                if p.M & (1 << keyBitMap[grid[p.R][p.C]]) == 0 {
                    continue
                }
            }

            if grid[p.R][p.C] >= 'a' && grid[p.R][p.C] <= 'z' {
                p.M = p.M | (1 << keyBitMap[grid[p.R][p.C]])
            }

            if p.M == (1 << klCnt - 1) {
                return result
            }

            if !dp[p.R][p.C][p.M] {
                dp[p.R][p.C][p.M] = true

                q.Push(Point{p.R-1, p.C, p.M})
                q.Push(Point{p.R+1, p.C, p.M})
                q.Push(Point{p.R, p.C-1, p.M})
                q.Push(Point{p.R, p.C+1, p.M})
            }
        }

        result++
    }

    return -1
}

