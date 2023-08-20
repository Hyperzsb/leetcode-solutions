class RecentCounter {
    private List<Integer> requests;
    
    public RecentCounter() {
        requests = new ArrayList<>();
    }
    
    public int ping(int t) {
        requests.add(t);

        int s = 0, e = requests.size() - 1;
        while (s < e) {
            int h = (s + e) / 2;

            if (requests.get(h) >= t - 3000) {
                e = h;
            } else {
                s = h + 1;
            }
        }

        return requests.size() - s;
    }
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter obj = new RecentCounter();
 * int param_1 = obj.ping(t);
 */

