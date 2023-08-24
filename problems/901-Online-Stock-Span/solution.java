class StockSpanner {
    private List<Integer> days;
    private List<Integer> prices;

    public StockSpanner() {
        days = new ArrayList<>();
        days.add(0);
        prices = new ArrayList<>();
        prices.add(Integer.MAX_VALUE);
    }
    
    public int next(int price) {
        while (prices.get(days.get(days.size() - 1)) <= price) {
            days.remove(days.size() - 1);
        }

        int result = prices.size() - days.get(days.size()-1);

        days.add(prices.size());
        prices.add(price);

        return result;
    }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * StockSpanner obj = new StockSpanner();
 * int param_1 = obj.next(price);
 */

