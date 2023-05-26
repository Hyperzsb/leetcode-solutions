class Solution {
    public boolean isPalindrome(String s) {
        int left = 0, right = s.length()-1;

        boolean result = true;
        while (left <= right) {
            if (s.charAt(left) == s.charAt(right)) {
                left++;
                right--;
                continue;
            } else {
                result = false;
                break;
            }
        }

        return result;
    }
    public boolean validPalindrome(String s) {
        int left = 0, right = s.length()-1;
        
        boolean flag = false, result = true;
        while (left <= right) {
            if (s.charAt(left) == s.charAt(right)) {
                left++;
                right--;
                continue;
            }

            if (!isPalindrome(s.substring(left+1, right+1)) && 
                !isPalindrome(s.substring(left, right))) {
                result = false;
            }

            break;
        }

        return result;
    }
}

