/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* sortedSquares(int* nums, int numsSize, int* returnSize) {
    int negativeNumsSize = 0;
    while(negativeNumsSize < numsSize && nums[negativeNumsSize] < 0) {
        negativeNumsSize++;
    }
    
    int *negativeNums = (int*)malloc((negativeNumsSize + 1) * sizeof(int));
    int *positiveNums = (int*)malloc((numsSize - negativeNumsSize + 1) * sizeof(int));
    for(int i = 0; i < negativeNumsSize; i++) {
        negativeNums[negativeNumsSize - 1 - i] = nums[i] * nums[i];
    }
    for(int i = negativeNumsSize; i < numsSize; i++) {
        positiveNums[i - negativeNumsSize] = nums[i] * nums[i];
    }
    
    int *returnNums = (int*)malloc((numsSize + 1) * sizeof(int));
    int index = 0, negativeIndex = 0, positiveIndex = 0;
    while(negativeIndex < negativeNumsSize && positiveIndex < (numsSize - negativeNumsSize)) {
        if(negativeNums[negativeIndex] <= positiveNums[positiveIndex]) {
            returnNums[index] = negativeNums[negativeIndex];
            negativeIndex++;
            index++;
        } else {
            returnNums[index] = positiveNums[positiveIndex];
            positiveIndex++;
            index++;
        }
    }
    while(negativeIndex < negativeNumsSize) {
        returnNums[index] = negativeNums[negativeIndex];
        negativeIndex++;
        index++;
    }
    while(positiveIndex < (numsSize - negativeNumsSize)) {
        returnNums[index] = positiveNums[positiveIndex];
        positiveIndex++;
        index++;
    }
    
    for(int i = 0; i < numsSize; i++) {
        printf("%d ", returnNums[i]);
    }
    
    free(negativeNums);
    free(positiveNums);
        
    *returnSize = numsSize;
    
    return returnNums;
}

