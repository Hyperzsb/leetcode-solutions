class Solution {
    public boolean isSelfDividing(int num) {
        int temp = num;
        while (temp > 0) {
            int divisor = temp % 10;
            temp /= 10;

            if (divisor == 0 || num % divisor != 0) {
                return false;
            }
        }

        return true;
    }

    public List<Integer> selfDividingNumbers(int left, int right) {
        List<Integer> result = new ArrayList<>();
        for (int i = left; i <= right; i++) {
            if (isSelfDividing(i)) {
                result.add(i);
            }
        }

        return result;
    }
}

