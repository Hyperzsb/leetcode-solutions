int minCostClimbingStairs(int* cost, int costSize){
    int dp[1010];
    memset(dp, 0, 1010 * sizeof(int));
    
    dp[costSize - 1] = cost[costSize - 1];
    dp[costSize - 2] = cost[costSize - 2];
    
    for(int i = costSize - 3; i >= 0; i--) {
        dp[i] = dp[i + 1] < dp[i + 2] ? dp[i + 1] + cost[i] : dp[i + 2] + cost[i];
    }
    
    return dp[0] < dp[1] ? dp[0] : dp[1];
}

