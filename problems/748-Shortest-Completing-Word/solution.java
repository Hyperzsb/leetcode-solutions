class Solution {
    public boolean isCompleting(String word, Map<Character, Integer> cm) {
        char[] nca = word.toCharArray();
        Map<Character, Integer> ncm = new HashMap<>();
        for (char c : nca) {
            Integer cnt = ncm.get(c);
            if (cnt == null) {
                ncm.put(c, 1);
            } else {
                ncm.put(c, cnt + 1);
            }
        }

        for (char c : cm.keySet()) {
            Integer cnt = ncm.get(c);

            if (cnt == null) {
                return false;
            }

            if (cnt < cm.get(c)) {
                return false;
            }
        }

        return true;
    }

    public String shortestCompletingWord(String licensePlate, String[] words) {
        char[] ca = licensePlate.toCharArray();
        Map<Character, Integer> cm = new HashMap<>();
        for (char c : ca) {
            if (c >= 'a' && c <= 'z') {
                Integer cnt = cm.get(c);
                if (cnt == null) {
                    cm.put(c, 1);
                } else {
                    cm.put(c, cnt + 1);
                }
            } else if (c >= 'A' && c <= 'Z') {
                Integer cnt = cm.get((char)(c - 'A' + 'a'));
                if (cnt == null) {
                    cm.put((char)(c - 'A' + 'a'), 1);
                } else {
                    cm.put((char)(c - 'A' + 'a'), cnt + 1);
                }
            } else {
                continue;
            }
        }

        Arrays.sort(words, (a, b) -> a.length() - b.length());

        String result = "";
        for (String w : words) {
            if (isCompleting(w, cm)) {
                result = w;
                break;
            }
        }

        return result;
    }
}

