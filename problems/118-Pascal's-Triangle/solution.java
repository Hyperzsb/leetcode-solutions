class Solution {
    public List<List<Integer>> generate(int numRows) {
        List<List<Integer>> result = new LinkedList<>();

        result.add(new LinkedList<>());
        result.get(0).add(1);

        for (int i = 1; i < numRows; i++) {
            result.add(new LinkedList<>());
            result.get(i).add(1);

            for (int j = 1; j < i; j++) {
                result.get(i).add(result.get(i-1).get(j-1) + result.get(i-1).get(j));
            }

            result.get(i).add(1);
        }

        return result;
    }
}

