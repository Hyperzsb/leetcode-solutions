class Solution {
    public int countBinarySubstrings(String s) {
        ArrayList<Integer> digits = new ArrayList<>();
        int cnt = 1;
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) != s.charAt(i-1)) {
                digits.add(cnt);
                cnt = 1;
            } else {
                cnt++;
            }
        }
        digits.add(cnt);

        int result = 0;
        for (int i = 1; i < digits.size(); i++) {
            result += Math.min(digits.get(i-1), digits.get(i));
        }

        return result;
    }
}

