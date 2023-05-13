class Solution {
    public int findLHS(int[] nums) {
        Arrays.sort(nums);
        int prevIdx = -1, currIdx = 0, result = 0;
        
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == nums[currIdx]) {
                continue;
            }

            if (prevIdx == -1) {
                prevIdx = currIdx;
                currIdx = i;
                continue;
            }

            if (nums[currIdx] - nums[prevIdx] == 1) {
                result = result > i - prevIdx ? result : i - prevIdx;
            }

            prevIdx = currIdx;
            currIdx = i;
        }

        if (prevIdx != -1 && nums[currIdx] - nums[prevIdx] == 1) {
            result = result > nums.length - prevIdx ? result : nums.length - prevIdx;
        }

        return result;
    }
}

