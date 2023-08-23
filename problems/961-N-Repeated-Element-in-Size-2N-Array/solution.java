class Solution {
    public int repeatedNTimes(int[] nums) {
        Set<Integer> cnts = new HashSet<>();
        for (int num : nums) {
            if (cnts.contains(num)) {
                return num;
            } else {
                cnts.add(num);
            }
        }

        return 0;
    }
}

