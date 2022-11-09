int pivotIndex(int* nums, int numsSize) {
    int *runningSum = (int *)malloc(numsSize * sizeof(int));
    runningSum[0] = nums[0];
    
    for(int i = 1; i < numsSize; i++) {
        runningSum[i] = runningSum[i - 1] + nums[i];
    }
    
    
    if(runningSum[numsSize - 1] - runningSum[0] == 0) {
        return 0;
    }
    
    for(int i = 1; i < numsSize - 1; i++) {
        if(runningSum[i - 1] == runningSum[numsSize - 1] - runningSum[i]) {
            return i;
        }
    }
    
    if(runningSum[numsSize - 2] == 0) {
        return numsSize - 1;
    }
    
    return -1;
}

