class Solution {
    private int count(int n) {
        int result = 0;
        while (n > 0) {
            if ((n & 1) == 1) {
                result++;
            }

            n >>= 1;
        }

        return result;
    }

    public int[] countBits(int n) {
        int[] result = new int[n + 1];
        for (int i = 0; i <= n; i++) {
            result[i] = count(i);
        }

        return result;
    }
}

