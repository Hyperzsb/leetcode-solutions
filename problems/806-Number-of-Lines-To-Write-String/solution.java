class Solution {
    public int[] numberOfLines(int[] widths, String s) {
        int line = 1, width = 0;
        for (int i = 0; i < s.length();) {
            if (width + widths[s.charAt(i) - 'a'] > 100) {
                line++;
                width = 0;
            } else {
                width += widths[s.charAt(i) - 'a'];
                i++;
            }
        }

        return new int[]{line, width};
    }
}

