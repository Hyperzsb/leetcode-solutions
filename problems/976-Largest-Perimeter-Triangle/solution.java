class Solution {
    public int largestPerimeter(int[] nums) {
        Arrays.sort(nums);
        for (int i = nums.length - 1; i >= 0; i--) {
            for (int j = i - 1; j >= 0; j--) {
                for (int k = j - 1; k >= 0; k--) {
                    if (nums[k] + nums[j] > nums[i]) {
                        return nums[i] + nums[j] + nums[k];
                    }
                }
            }
        }

        return 0;
    }
}

