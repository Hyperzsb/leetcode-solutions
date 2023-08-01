class Solution {
    public boolean isMonotonic(int[] nums) {
        if (nums.length < 2) {
            return true;
        }

        int mono = 0;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i-1] == nums[i]) {
                continue;
            }

            if (nums[i-1] < nums[i]) {
                if (mono == 0) {
                    mono = 1;
                } else if (mono == 1) {
                    continue;
                } else {
                    return false;
                }
            } else {
                if (mono == 0) {
                    mono = -1;
                } else if (mono == -1) {
                    continue;
                } else {
                    return false;
                }
            }
        }

        return true;
    }
}

