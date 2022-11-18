/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* runningSum(int* nums, int numsSize, int* returnSize){
    int *runningSum = (int *)malloc(numsSize * sizeof(int));
    *returnSize = numsSize;
    runningSum[0] = nums[0];
    for(int i = 1; i < numsSize; i++) 
        runningSum[i] = runningSum[i - 1] + nums[i];
    return runningSum;
}

