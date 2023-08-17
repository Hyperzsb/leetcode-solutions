func change(amount int, coins []int) int {
    if amount == 0 {
        return 1
    }
    
    sort.Slice(coins, func(a, b int) bool {
        return coins[a] < coins[b]
    })

    dp := make([][]int, amount+1)
    for i := range dp {
        dp[i] = make([]int, len(coins))
    }

    for i := 1; i <= amount; i++ {
        for j := 0; j < len(coins) && coins[j] <= i; j++ {
            if coins[j] == i {
                dp[i][j] = 1
                continue
            }

            for k := 0; k <= j; k++ {
                dp[i][j] += dp[i-coins[j]][k]
            }
        }
    }

    result := 0
    for j := range dp[amount] {
        result += dp[amount][j]
    }

    return result
}

