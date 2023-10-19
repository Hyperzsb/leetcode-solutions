class Solution {
    public List<Integer> getRow(int rowIndex) {
        List<Integer> result = new LinkedList<>();
        result.add(1);

        while (rowIndex > 0) {
            List<Integer> temp = new LinkedList<>();

            temp.add(1);
            for (int i = 1; i < result.size(); i++) {
                temp.add(result.get(i - 1) + result.get(i));
            }
            temp.add(1);

            result = temp;

            rowIndex--;
        }

        return result;
    }
}

