class Solution {
    public String longestPalindrome(String s) {
        List<List<Integer>> checkList = new ArrayList<>();
        for (int i = 0; i < s.length(); i++) {
            checkList.add(new ArrayList<>());
            checkList.get(i).add(0);
            checkList.get(i).add(1);
        }

        int maxLen = 1;
        String maxStr = s.substring(0, 1);
        for (int i = 1; i < s.length(); i++) {
            for (int j = 0; j < checkList.get(i - 1).size(); j++) {
                if (i - 1 - checkList.get(i - 1).get(j) >= 0 && s.charAt(i - 1 - checkList.get(i - 1).get(j)) == s.charAt(i)) {
                    checkList.get(i).add(checkList.get(i - 1).get(j) + 2);

                    if (maxLen < checkList.get(i - 1).get(j) + 2) {
                        maxLen = checkList.get(i - 1).get(j) + 2;
                        maxStr = s.substring(i - 1 - checkList.get(i - 1).get(j), i + 1);
                    }
                }
            }
        }

        return maxStr;
    }
}

