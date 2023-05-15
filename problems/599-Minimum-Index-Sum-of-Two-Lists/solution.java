class Solution {
    public String[] findRestaurant(String[] list1, String[] list2) {
        LinkedList<String> result = new LinkedList<>();
        int minIdxSum = list1.length + list2.length;

        for (int i = 0; i < list1.length; i++) {
            if (i > minIdxSum) {
                break;
            }

            for (int j = 0; j < list2.length; j++) {
                if (i + j > minIdxSum) {
                    break;
                }

                if (list1[i].equals(list2[j])) {
                    if (i + j == minIdxSum) {
                        result.add(list1[i]);
                    } else {
                        minIdxSum = i + j;
                        result.clear();
                        result.add(list1[i]);
                    }
                }
            }
        }

        return result.toArray(new String[result.size()]);
    }
}

