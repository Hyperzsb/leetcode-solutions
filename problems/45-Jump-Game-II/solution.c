int jump(int* nums, int numsSize){
    int *dp = (int *)malloc(numsSize * sizeof(int));
    for(int i = 0; i < numsSize; i++) {
        dp[i] = 100000;
    }
    
    dp[numsSize - 1] = 0;
    
    for(int i = numsSize - 2; i >= 0; i--) {
        for(int j = 1; j <= nums[i] && i + j <= numsSize - 1; j++) {
            if(dp[i] > dp[i + j] + 1) {
                dp[i] = dp[i + j] + 1;
            }
        }
    }
    
    int min = dp[0];
    
    free(dp);
    
    return min;
}

