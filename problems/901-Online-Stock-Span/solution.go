type StockSpanner struct {
    Prices []int
    Spans []int
}


func Constructor() StockSpanner {
    return StockSpanner{
        Prices: make([]int, 0),
        Spans: make([]int, 0),
    }
}


func (this *StockSpanner) Next(price int) int {
    this.Prices = append(this.Prices, price)

    if len(this.Prices) == 1 {
        this.Spans = append(this.Spans, 1)
    } else {
        curr, prev := len(this.Prices)-1, len(this.Prices)-2
        for {
            if this.Prices[prev] <= this.Prices[curr] {
                prev -= this.Spans[prev]
            } else {
                break
            }

            if prev < 0 {
                break
            }
        }

        this.Spans = append(this.Spans, curr-prev)
    }

    return this.Spans[len(this.Spans)-1]
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

