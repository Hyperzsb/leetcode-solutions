class Solution {
    public int numJewelsInStones(String jewels, String stones) {
        HashSet<Character> cs = new HashSet<>();
        for (char c : jewels.toCharArray()) {
            cs.add(c);
        }

        int result = 0;
        for (char c : stones.toCharArray()) {
            if (cs.contains(c)) {
                result++;
            }
        }

        return result;
    }
}

