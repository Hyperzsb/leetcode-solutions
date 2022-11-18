typedef struct {
    int days;
    int *prices;
    int *spans;
} StockSpanner;


StockSpanner* stockSpannerCreate() {
    StockSpanner *spanner = (StockSpanner *)malloc(sizeof(StockSpanner));
    spanner->days = 0;
    spanner->prices = (int *)malloc(10005 * sizeof(int));
    spanner->spans = (int *)malloc(10005 * sizeof(int));
    spanner->prices[0] = 1000000;
    spanner->spans[0] = 1;
    return spanner;
}

int stockSpannerNext(StockSpanner* obj, int price) {
    obj->days += 1;
    int today = obj->days;
    obj->prices[today] = price;
    int cursor = obj->days - 1;
    while(true) {
        if(obj->prices[cursor] > obj->prices[today]) {
            obj->spans[today] = today - cursor;
            break;
        } else {
            cursor -= obj->spans[cursor];
        }
    }
    
    return obj->spans[today];
}

void stockSpannerFree(StockSpanner* obj) {
    free(obj->prices);
    free(obj->spans);
    free(obj);
}

/**
 * Your StockSpanner struct will be instantiated and called as such:
 * StockSpanner* obj = stockSpannerCreate();
 * int param_1 = stockSpannerNext(obj, price);
 
 * stockSpannerFree(obj);
*/

