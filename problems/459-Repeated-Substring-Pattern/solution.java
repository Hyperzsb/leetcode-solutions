class Solution {
    public boolean repeatedSubstringPattern(String s) {
        for (int i = 1; i < s.length(); i++) {
            if (s.length() % i != 0) {
                continue;
            }

            boolean flag = true;
            for (int j = 1; j * i < s.length(); j++) {
                if (!s.substring(0, i).equals(s.substring(j * i, (j + 1) * i))) {
                    flag = false;
                    break;
                }
            }

            if (flag) {
                return true;
            }
        }

        return false;
    }
}

