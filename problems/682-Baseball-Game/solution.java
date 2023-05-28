class Solution {
    public int calPoints(String[] operations) {
        ArrayList<Integer> records = new ArrayList<>();

        for (int i = 0; i < operations.length; i++) {
            if (operations[i].equals("+")) {
                records.add(records.get(records.size() - 1) + records.get(records.size() - 2));
            } else if (operations[i].equals("D")) {
                records.add(records.get(records.size() - 1) * 2);
            } else if (operations[i].equals("C")) {
                records.remove(records.size() - 1);
            } else {
                records.add(Integer.parseInt(operations[i]));
            }
        }

        int sum = 0;
        for (int record : records) {
            sum += record;
        }

        return sum;
    }
}

