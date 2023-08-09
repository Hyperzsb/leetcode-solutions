class Solution {
    public boolean buddyStrings(String s, String goal) {
        if (s.length() != goal.length()) {
            return false;
        }

        List<Integer> diffIdxes = new LinkedList<>();
        Map<Character, Integer> charCnt = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            if (charCnt.containsKey(s.charAt(i))) {
                charCnt.put(s.charAt(i), charCnt.get(s.charAt(i)) + 1);
            } else {
                charCnt.put(s.charAt(i), 1);
            }

            if (s.charAt(i) == goal.charAt(i)) {
                continue;
            }

            if (diffIdxes.size() == 2) {
                return false;
            }

            diffIdxes.add(i);
        }

        if (diffIdxes.size() == 0) {
            for (Map.Entry<Character, Integer> entry : charCnt.entrySet()) {
                if (entry.getValue() >= 2) {
                    return true;
                }
            }

            return false;
        } else if (diffIdxes.size() == 1) {
            return false;
        } else {
            if (s.charAt(diffIdxes.get(0)) == goal.charAt(diffIdxes.get(1)) && s.charAt(diffIdxes.get(1)) == goal.charAt(diffIdxes.get(0))) {
                return true;
            } else {
                return false;
            }
        }
    }
}

