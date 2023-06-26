class Solution {
    public String toGoatLatin(String sentence) {
        String[] words = sentence.split(" ");
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            StringBuilder latin = new StringBuilder(w);

            if (!(w.charAt(0) == 'a' || w.charAt(0) == 'A' ||
                w.charAt(0) == 'e' || w.charAt(0) == 'E' ||
                w.charAt(0) == 'i' || w.charAt(0) == 'I' ||
                w.charAt(0) == 'o' || w.charAt(0) == 'O' ||
                w.charAt(0) == 'u' || w.charAt(0) == 'U')) {
                char c = latin.charAt(0);
                latin.deleteCharAt(0);
                latin.append(c);
            }

            latin.append("ma");
            for (int j = 0; j <= i; j++) {
                latin.append('a');
            }

            words[i] = latin.toString();
        }

        String result = new String();
        for (String w : words) {
            result += w + " ";
        }

        return result.substring(0, result.length()-1);
    }
}

