class Solution {
    public List<String> summaryRanges(int[] nums) {
        ArrayList<String> result = new ArrayList<>();

        if (nums.length == 0) {
            return result;
        }

        int start = nums[0], prev = nums[0];
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == prev + 1) {
                prev++;
                continue;
            }

            if (start == prev) {
                result.add(String.format("%d", start));
            } else {
                result.add(String.format("%d->%d", start, prev));
            }

            start = nums[i];
            prev = nums[i];
        }

        if (start == prev) {
            result.add(String.format("%d", start));
        } else {
            result.add(String.format("%d->%d", start, prev));
        }

        return result;
    }
}

