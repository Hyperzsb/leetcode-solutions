class Solution {
    public int findLengthOfLCIS(int[] nums) {
        int result = 1, left = 0;
        for (int right = 1; right < nums.length; right++) {
            if (nums[right] <= nums[right-1]) {
                result = result > right - left ? result : right - left;
                left = right;
            }
        }

        result = result > nums.length - left ? result : nums.length - left;

        return result;
    }
}

