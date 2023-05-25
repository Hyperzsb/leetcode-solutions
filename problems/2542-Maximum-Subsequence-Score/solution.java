class Solution {
    class Number {
        int add;
        int mul;

        public Number(int add, int mul) {
            this.add = add;
            this.mul = mul;
        }

        @Override
        public String toString() {
            return "[" + add + ", " + mul + "]";
        }
    }

    class Sortbymul implements Comparator<Number> {
        public int compare(Number a, Number b) {
            return b.mul - a.mul;
        }
    }

    public long maxScore(int[] nums1, int[] nums2, int k) {
        ArrayList<Number> nums = new ArrayList<>();
        for (int i = 0; i < nums1.length; i++) {
            nums.add(new Number(nums1[i], nums2[i]));
        }

        Collections.sort(nums, (a, b) -> b.mul - a.mul);

        PriorityQueue<Integer> pq = new PriorityQueue<>(k, (a, b) -> a - b);

        long result = 0, sum = 0;
        for (Number num : nums) {
            pq.add(num.add);
            sum += num.add;

            if (pq.size() > k) {
                sum -= pq.poll();
            }

            if (pq.size() == k) {
                result = result > sum * num.mul ? result : sum * num.mul;
            }
        }

        return result;
    }
}

