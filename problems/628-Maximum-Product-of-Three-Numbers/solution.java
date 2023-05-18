class Solution {
    public int maximumProduct(int[] nums) {
        Arrays.sort(nums);

        int result = nums[0] * nums[1] * nums[nums.length - 1];
        if (result < nums[nums.length - 1] * nums[nums.length - 2] * nums[nums.length - 3]) {
            result = nums[nums.length - 1] * nums[nums.length - 2] * nums[nums.length - 3];
        }

        return result;
    }
}

