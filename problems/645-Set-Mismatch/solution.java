class Solution {
    public int[] findErrorNums(int[] nums) {
        int duplicate = 0, loss = 0;
        int expectedSum = (1 + nums.length) * nums.length / 2;
        int targetSum = 0;

        HashSet<Integer> set = new HashSet<>();
        for (int n : nums) {
            if (set.contains(n)) {
                duplicate = n;
            } else {
                set.add(n);
            }

            targetSum += n;
        }

        loss = expectedSum - targetSum + duplicate;

        return new int[]{duplicate, loss};
    }
}

