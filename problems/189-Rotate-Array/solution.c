void rotate(int* nums, int numsSize, int k){  
    int *tempNums = (int *)malloc(numsSize * sizeof(int));
    
    int index = 0;
    for(int i = 0; i < numsSize; i++) {
        index = (numsSize - k + i)%numsSize >= 0 ? (numsSize - k + i)%numsSize : numsSize + (numsSize - k + i)%numsSize;
        tempNums[i] = nums[index];
    }
    
    for(int i = 0; i < numsSize; i++) {
        nums[i] = tempNums[i];
    }
}

