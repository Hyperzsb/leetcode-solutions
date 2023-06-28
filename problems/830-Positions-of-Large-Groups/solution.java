class Solution {
    public List<List<Integer>> largeGroupPositions(String s) {
        ArrayList<List<Integer>> result = new ArrayList<>();

        int idx = 0;
        char c = s.charAt(0);
        for (int i = 1; i < s.length(); i++) {
            if (c == s.charAt(i)) {
                continue;
            }

            if (i - idx >= 3) {
                result.add(new ArrayList(Arrays.asList(idx, i - 1)));
            }

            idx = i;
            c = s.charAt(i);
        }

        if (s.length() - idx >= 3) {
            result.add(new ArrayList(Arrays.asList(idx, s.length() - 1)));
        }

        return result;
    }
}

