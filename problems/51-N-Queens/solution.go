func convert(board [][]int) []string {
    result := make([]string, len(board))
    
    for i := range board {
        row := ""
        for j := range board[i] {
            if board[i][j] == -1 {
                row += "Q"
            } else {
                row += "."
            }
        }
        result[i] = row
    }

    return result
}

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

func nqueens(idx, n int, board [][]int, result *[][]string) {
    if idx == n {
        newResult := convert(board)
        *result = append(*result, newResult)

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

func solveNQueens(n int) [][]string {
    board := make([][]int, n)
    for i := range board {
        board[i] = make([]int, n)
    }

    result := make([][]string, 0)
    nqueens(0, n, board, &result)

    return result
}

