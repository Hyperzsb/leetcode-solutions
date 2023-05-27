class Solution {
    public boolean hasAlternatingBits(int n) {
        int flag = n & 1;
        n >>= 1;
        while (n > 0) {
            if ((flag ^ (n & 1)) == 0) {
                return false;
            }
            flag = n & 1;
            n >>= 1;
        }

        return true;
    }
}

