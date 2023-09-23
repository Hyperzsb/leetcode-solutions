func numRookCaptures(board [][]byte) int {
    x, y := 0, 0
    for i := range board {
        flag := false
        for j := range board[i] {
            if board[i][j] == 'R' {
                x, y = i, j
                break
            }
        }

        if flag {
            break
        }
    }

    result := 0

    for i := x - 1; i >= 0; i-- {
        if board[i][y] == 'p' {
            result++
            break
        } else if board[i][y] == 'B' {
            break
        } else {
            continue
        }
    }

    for i := x + 1; i < len(board); i++ {
        if board[i][y] == 'p' {
            result++
            break
        } else if board[i][y] == 'B' {
            break
        } else {
            continue
        }
    } 

    for i := y - 1; i >= 0; i-- {
        if board[x][i] == 'p' {
            result++
            break
        } else if board[x][i] == 'B' {
            break
        } else {
            continue
        }
    }

    for i := y + 1; i < len(board[0]); i++ {
        if board[x][i] == 'p' {
            result++
            break
        } else if board[x][i] == 'B' {
            break
        } else {
            continue
        }
    } 

    return result
}

