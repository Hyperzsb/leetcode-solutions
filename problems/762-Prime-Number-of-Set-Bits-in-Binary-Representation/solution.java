class Solution {
    public boolean isPrime(int n) {
        if (n == 1) {
            return false;
        }

        for (int i = 2; i * i <= n; i++) {
            if (n % i == 0) {
                return false;
            }
        }

        return true;
    }

    public boolean hasPrimeBits(int n) {
        int cnt = 0;
        while(n > 0) {
            int temp = n & 1;
            if (temp == 1) {
                cnt++;
            }

            n >>= 1;
        }

        return isPrime(cnt);
    }

    public int countPrimeSetBits(int left, int right) {
        int result = 0;
        for (int i = left; i <= right; i++) {
            if (hasPrimeBits(i)) {
                result++;
            }
        }

        return result;
    }
}

