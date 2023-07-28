class Solution {
    public String[] uncommonFromSentences(String s1, String s2) {
        String[] words1 = s1.split(" ");

        Map<String, Integer> wordCnts1 = new HashMap<>();
        for (String word : words1) {
            if (wordCnts1.containsKey(word)) {
                wordCnts1.put(word, wordCnts1.get(word) + 1);
            } else {
                wordCnts1.put(word, 1);
            }
        }

        String[] words2 = s2.split(" ");

        Map<String, Integer> wordCnts2 = new HashMap<>();
        for (String word : words2) {
            if (wordCnts2.containsKey(word)) {
                wordCnts2.put(word, wordCnts2.get(word) + 1);
            } else {
                wordCnts2.put(word, 1);
            }
        }

        List<String> result = new LinkedList<>();
        
        for (Map.Entry<String, Integer> e : wordCnts1.entrySet()) {
            if (e.getValue() > 1) {
                continue;
            }

            if (!wordCnts2.containsKey(e.getKey())) {
                result.add(e.getKey());
            } else {
                wordCnts2.remove(e.getKey());
            }
        }

        for (Map.Entry<String, Integer> e : wordCnts2.entrySet()) {
            if (e.getValue() > 1) {
                continue;
            }
            
            if (!wordCnts1.containsKey(e.getKey())) {
                result.add(e.getKey());
            } else {
                wordCnts1.remove(e.getKey());
            }
        }

        return result.toArray(new String[result.size()]);
    }
}

