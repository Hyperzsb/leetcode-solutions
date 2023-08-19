class Solution {
    public int findKthLargest(int[] nums, int k) {
        PriorityQueue<Integer> pq = new PriorityQueue<Integer>(k, (a, b) -> a - b);

        for (int num : nums) {
            if (pq.size() < k) {
                pq.add(num);
                continue;
            }

            if (pq.peek() < num) {
                pq.poll();
                pq.add(num);
            }
        }

        return pq.peek();
    }
}

