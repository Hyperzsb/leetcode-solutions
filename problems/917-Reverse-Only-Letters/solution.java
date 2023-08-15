class Solution {
    private boolean isLetter(char c) {
        if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
            return true;
        } else {
            return false;
        }
    }

    public String reverseOnlyLetters(String s) {
        int si = 0, ei = s.length()-1;

        StringBuilder sb = new StringBuilder(s);
        while (si < ei) {
            while (si <= s.length()-1 && !isLetter(sb.charAt(si))) {
                si++;
            }

            while (ei >= 0 && !isLetter(sb.charAt(ei))) {
                ei--;
            }

            if (si >= ei) {
                break;
            }

            char sc = sb.charAt(si), ec = sb.charAt(ei);
            sb.setCharAt(si, ec);
            sb.setCharAt(ei, sc);

            si++;
            ei--;
        }

        return sb.toString();
    }
}

