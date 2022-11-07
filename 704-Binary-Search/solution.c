int search(int* nums, int numsSize, int target){
    int start = 0, end = numsSize - 1, half = (start + end) / 2;
    while (start < end) {
        if (target == nums[half]) {
            return half;
        } else if (target < nums[half]) {
            end = half - 1;
        } else {
            start = half + 1;          
        }
        half =  (start + end) / 2;
    }
    if (target == nums[half]) {
        return half;
    } else {
        return -1;
    }
}
