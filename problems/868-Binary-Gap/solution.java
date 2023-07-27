class Solution {
    public int binaryGap(int n) {
        int result = 0, curr = 0, prev = -1;
        while (n > 0) {
            int bit = n & 1;

            if (bit == 1) {
                if (prev == -1) {
                    prev = curr;
                } else {
                    if (curr - prev > result) {
                        result = curr - prev;
                    }

                    prev = curr;
                }
            }
            
            n >>= 1;
            curr++;
        }

        return result;
    }
}

