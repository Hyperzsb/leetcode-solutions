func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    dp := make([][]int, len(obstacleGrid))
    for i := range dp {
        dp[i] = make([]int, len(obstacleGrid[i]))
    }

    if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
        return 0
    } 

    dp[0][0] = 1

    for i := range obstacleGrid {
        for j := range obstacleGrid[i] {
            if obstacleGrid[i][j] == 1 {
                continue
            }

            if i > 0 {
                dp[i][j] += dp[i-1][j]
            }

            if j > 0 {
                dp[i][j] += dp[i][j-1]
            }
        }
    }

    return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

