class Solution {
    public String rotate(String s) {
        StringBuilder sb = new StringBuilder(s);

        return sb.append(sb.charAt(0)).deleteCharAt(0).toString();
    }

    public boolean rotateString(String s, String goal) {
        for (int i = 0; i < s.length(); i++) {
            if (goal.equals(s)) {
                return true;
            }

            s = rotate(s);
        }

        return false;
    }
}

