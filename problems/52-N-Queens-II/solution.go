func add(r, c int, board [][]int) {
    for i := 1; i < len(board)-r; i++ {
        if c-i >= 0 {
            board[r+i][c-i]++
        }

        board[r+i][c]++

        if c+i <= len(board[r+i])-1 {
            board[r+i][c+i]++
        }
    }

    board[r][c] = -1
}

func remove(r, c int, board [][]int) {
    for i := 1; i < len(board)-r; i++ {
        if c-i >= 0 {
            board[r+i][c-i]--
        }

        board[r+i][c]--

        if c+i <= len(board[r+i])-1 {
            board[r+i][c+i]--
        }
    }

    board[r][c] = 0
}

func nqueens(idx, n int, board [][]int, result *int) {
    if idx == n {
        (*result)++

        return
    }

    for i := 0; i < n; i++ {
        if board[idx][i] == 0 {
            add(idx, i, board)
            nqueens(idx+1, n, board, result)
            remove(idx, i, board)
        }
    }
}

func totalNQueens(n int) int {
    board := make([][]int, n)
    for i := range board {
        board[i] = make([]int, n)
    }

    result := 0
    nqueens(0, n, board, &result)

    return result
}

