class Solution {
    public boolean judgeCircle(String moves) {
        int up = 0, down = 0, left = 0, right = 0;
        for (char dir : moves.toCharArray()) {
            switch (dir) {
            case 'U':
                up++;
                break;
            case 'D':
                down++;
                break;
            case 'L':
                left++;
                break;
            case 'R':
                right++;
                break;
            }
        }

        if (up == down && left == right) {
            return true;
        } else {
            return false;
        }
    }
}

