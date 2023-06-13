class Solution {
    private static String[] morseCodes = {
        ".-",
        "-...",
        "-.-.",
        "-..",
        ".",
        "..-.",
        "--.",
        "....",
        "..",
        ".---",
        "-.-",
        ".-..",
        "--",
        "-.",
        "---",
        ".--.",
        "--.-",
        ".-.",
        "...",
        "-",
        "..-",
        "...-",
        ".--",
        "-..-",
        "-.--",
        "--.."
    };

    private String toMorseCode(String word) {
        StringBuilder sb = new StringBuilder();
        for (char c : word.toCharArray()) {
            sb.append(morseCodes[c - 'a']);
        }

        return sb.toString();
    }

    public int uniqueMorseRepresentations(String[] words) {
        HashSet<String> hs = new HashSet<>();
        for (String w : words) {
            String mc = toMorseCode(w);
            if (!hs.contains(mc)) {
                hs.add(mc);
            }
        }

        return hs.size();
    }
}

