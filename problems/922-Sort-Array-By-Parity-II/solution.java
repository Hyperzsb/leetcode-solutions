class Solution {
    private int findNext(int[] nums, int idx, boolean mode) {
        if (mode) {
            for (int i = idx; i < nums.length; i += 2) {
                if (nums[i] % 2 == 0) {
                    return i;
                }
            }
        } else {
            for (int i = idx; i < nums.length; i += 2) {
                if (nums[i] % 2 != 0) {
                    return i;
                }
            }
        }

        return -1;
    }

    public int[] sortArrayByParityII(int[] nums) {
        for (int i = 0; i < nums.length; i++) {
            if (i % 2 == 0) {
                if (nums[i] % 2 == 0) {
                    continue;
                }

                int evenIdx = findNext(nums, i + 1, true);

                nums[i] ^= nums[evenIdx];
                nums[evenIdx] ^= nums[i];
                nums[i] ^= nums[evenIdx];
            } else {
                if (nums[i] % 2 != 0) {
                    continue;
                }

                int oddIdx = findNext(nums, i + 1, false);

                nums[i] ^= nums[oddIdx];
                nums[oddIdx] ^= nums[i];
                nums[i] ^= nums[oddIdx];
            }
        }

        return nums;
    }
}

