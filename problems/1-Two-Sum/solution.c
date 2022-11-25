/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int cmp(int *a, int *b) {
    return *a - *b;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    *returnSize = 2;
    int *result = (int *)malloc(2 * sizeof(int));
    
    int *copies = (int *)malloc(numsSize * sizeof(int));
    for(int i = 0; i < numsSize; i++) {
        copies[i] = nums[i];
    }
    
    qsort(copies, numsSize, sizeof(int), cmp);
    
    int start = 0, end = numsSize - 1;
    while(start < end) {
        if(copies[start] + copies[end] == target) {
            break;
        } else if(copies[start] + copies[end] < target) {
            start++;
        } else {
            end--;
        }
    }
    
    int flag = 0;
    for(int i = 0; i < numsSize; i++) {
        if(nums[i] == copies[start] || nums[i] == copies[end]) {
            result[flag++] = i;
            if(flag == 2) {
                break;
            }
        }
    }
    
    free(copies);
    
    return result;
}

