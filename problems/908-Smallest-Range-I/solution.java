class Solution {
    public int smallestRangeI(int[] nums, int k) {
        int min = Integer.MAX_VALUE, max = Integer.MIN_VALUE;
        for (int num : nums) {
            if (min > num) {
                min = num;
            }

            if (max < num) {
                max = num;
            }
        }

        return Math.max(max - min - 2 * k, 0);
    }
}

