/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    *returnSize = 2;
    int *result = (int *)malloc(2 * sizeof(int));
    
    int start = 0, end = numsSize - 1;
    while(start < end) {
        if(nums[start] + nums[end] == target) {
            break;
        } else if(nums[start] + nums[end] < target) {
            start++;
        } else {
            end--;
        }
    }
    
    result[0] = start + 1;
    result[1] = end + 1;
    
    return result;
}

