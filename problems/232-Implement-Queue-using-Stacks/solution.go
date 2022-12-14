type MyQueue struct {
	s []int
}

func Constructor() MyQueue {
	return MyQueue{make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.s = append(this.s, x)
}

func (this *MyQueue) Pop() int {
	top := this.s[0]
	if len(this.s) > 1 {
		this.s = this.s[1:]
	} else {
		this.s = this.s[:0]
	}
	return top
}

func (this *MyQueue) Peek() int {
	return this.s[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.s) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

