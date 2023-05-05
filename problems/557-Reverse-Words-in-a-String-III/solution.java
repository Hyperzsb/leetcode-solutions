class Solution {
    public String reverseWords(String s) {
        String[] words = s.split(" ");
        for (int i = 0; i < words.length; i++) {
            char[] chars = words[i].toCharArray();
            for (int j = 0; j < chars.length / 2; j++) {
                chars[j] ^= chars[chars.length-1-j];
                chars[chars.length-1-j] ^= chars[j];
                chars[j] ^= chars[chars.length-1-j];
            }
            words[i] = new String(chars);
        }

        String result = String.join(" ", words);
        return result;
    }
}

