class Solution {
    public int[] shortestToChar(String s, char c) {
        int[] result = new int[s.length()];
        for (int i = 0; i < s.length(); i++) {
            int left = s.length();
            for (int j = 0; j <= i; j++) {
                if (s.charAt(i - j) == c) {
                    left = j;
                    break;
                }
            }

            int right = s.length();
            for (int j = 0; j < s.length() - i; j++) {
                if (s.charAt(i + j) == c) {
                    right = j;
                    break;
                }
            }

            result[i] = Math.min(left, right);
        }

        return result;
    }
}

