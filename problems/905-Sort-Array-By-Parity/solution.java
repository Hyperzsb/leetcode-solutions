class Solution {
    private int findNext(int[] nums, int idx, boolean flag) {
        int result = -1;

        if (flag) {
            for (int i = idx + 1; i < nums.length; i++) {
                if (nums[i] % 2 == 0) {
                    result = i;
                    break;
                }
            }
        } else {
            for (int i = idx + 1; i < nums.length; i++) {
                if (nums[i] % 2 != 0) {
                    result = i;
                    break;
                }
            }
        }

        return result;
    }

    public int[] sortArrayByParity(int[] nums) {
        int evenIdx = findNext(nums, -1, true), oddIdx = findNext(nums, -1, false);

        if (evenIdx == -1 || oddIdx == -1) {
            return nums;
        }

        while (evenIdx != -1) {
            if (evenIdx > oddIdx) {
                nums[evenIdx] ^= nums[oddIdx];
                nums[oddIdx] ^= nums[evenIdx];
                nums[evenIdx] ^= nums[oddIdx];

                evenIdx = findNext(nums, evenIdx, true);
                oddIdx = findNext(nums, oddIdx, false);
            } else {
                evenIdx = findNext(nums, oddIdx, true);
            }
        }

        return nums;
    }
}

