class Solution {
    public int findShortestSubArray(int[] nums) {
        Map<Integer, Integer> hm = new HashMap<>();
        for (int num : nums) {
            if (hm.get(num) == null) {
                hm.put(num, 1);
            } else {
                hm.put(num, hm.get(num) + 1);
            }
        }

        int maxCnt = 0;
        ArrayList<Integer> maxNums = new ArrayList<>();
        for (Map.Entry<Integer, Integer> e : hm.entrySet()) {
            if (e.getValue() == maxCnt) {
                if (!maxNums.contains(e.getKey())) {
                    maxNums.add(e.getKey());
                }
            }

            if (e.getValue() > maxCnt) {
                maxCnt = e.getValue();
                maxNums.clear();
                maxNums.add(e.getKey());
            }
        }

        int result = nums.length;
        for (int maxNum : maxNums) {
            int leftTrim = 0;
            while (nums[leftTrim] != maxNum) {
                leftTrim++;
            }

            int rightTrim = nums.length - 1;
            while (nums[rightTrim] != maxNum) {
                rightTrim--;
            }

            result = result < rightTrim - leftTrim + 1 ? result : rightTrim - leftTrim + 1;
        }

        return result;
    }
}

