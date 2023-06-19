func traverse(r, c, prev int, grid, dp [][]int) int {
    if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
        return 0
    }

    if grid[r][c] <= prev {
        return 0
    }

    if dp[r][c] != -1 {
        return dp[r][c]
    }

    cnt := 1
    cnt = (cnt + traverse(r-1, c, grid[r][c], grid, dp)) % (1e9 + 7)
    cnt = (cnt + traverse(r+1, c, grid[r][c], grid, dp)) % (1e9 + 7)
    cnt = (cnt + traverse(r, c-1, grid[r][c], grid, dp)) % (1e9 + 7)
    cnt = (cnt + traverse(r, c+1, grid[r][c], grid, dp)) % (1e9 + 7)

    dp[r][c] = cnt
    return cnt
}

func countPaths(grid [][]int) int {
    dp := make([][]int, len(grid))
    for i := range dp {
        dp[i] = make([]int, len(grid[i]))
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    result := 0
    for i := range dp {
        for j := range dp[i] {
            if dp[i][j] == -1 {
                result = (result + traverse(i, j, 0, grid, dp)) % (1e9 + 7)
            } else {
                result = (result + dp[i][j]) % (1e9 + 7)
            }
        }
    }

    return result
}

