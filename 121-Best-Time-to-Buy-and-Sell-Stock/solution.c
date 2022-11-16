int maxProfit(int* prices, int pricesSize){
    int maxPrice = 0;
    int maxProfit = 0;
    
    for(int i = pricesSize - 1; i >= 0; i --) {
        if(prices[i] > maxPrice) {
            maxPrice = prices[i];
        } else {
            maxProfit = maxProfit > maxPrice - prices[i] ? maxProfit : maxPrice - prices[i];
        }
    }
    
    return maxProfit;
}

