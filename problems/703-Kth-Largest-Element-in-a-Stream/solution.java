class KthLargest {
    int k;
    PriorityQueue<Integer> pq;

    public KthLargest(int k, int[] nums) {
        this.k = k;
        this.pq = new PriorityQueue<>(k);
        for (int num : nums) {
            this.add(num);
        }
    }
    
    public int add(int val) {
        if (this.pq.size() < k) {
            this.pq.add(val);
        } else {
            if (val > this.pq.peek()) {
                this.pq.poll();
                this.pq.add(val);
            }
        }

        return this.pq.peek();
    }
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest obj = new KthLargest(k, nums);
 * int param_1 = obj.add(val);
 */

