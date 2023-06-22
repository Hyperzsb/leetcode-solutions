class Solution {
    public String mostCommonWord(String paragraph, String[] banned) {
        HashSet<String> bhs = new HashSet<>();
        for (String bw : banned) {
            bhs.add(bw);
        }

        HashMap<String, Integer> whm = new HashMap<>();

        java.util.StringTokenizer pst = new java.util.StringTokenizer(paragraph, " !?',;.");

        int maxCnt = 0;
        String maxStr = "";
        while (pst.hasMoreTokens()) {
            String w = pst.nextToken().toLowerCase();
            if (bhs.contains(w)) {
                continue;
            }

            Integer cnt = whm.get(w);
            if (cnt == null) {
                whm.put(w, 1);

                if (maxCnt < 1) {
                    maxCnt = 1;
                    maxStr = w;
                }
            } else {
                whm.put(w, cnt + 1);

                if (maxCnt < cnt + 1) {
                    maxCnt = cnt + 1;
                    maxStr = w;
                }
            }
        }

        return maxStr;
    }
}

