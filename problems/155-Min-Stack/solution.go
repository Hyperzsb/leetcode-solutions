type MonoStack []int

func (ms *MonoStack) Push(num int) {
	*ms = append(*ms, num)
}

func (ms *MonoStack) Pop() int {
	top := (*ms)[len(*ms)-1]
	*ms = (*ms)[:len(*ms)-1]
	return top
}

func (ms *MonoStack) Top() int {
	return (*ms)[len(*ms)-1]
}

func (ms *MonoStack) Empty() bool {
	return len(*ms) == 0
}

type MinStack struct {
	s  []int
	ms MonoStack
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make(MonoStack, 0)}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	if this.ms.Empty() {
		this.ms.Push(len(this.s) - 1)
	} else {
		if this.s[this.ms.Top()] > val {
			this.ms.Push(len(this.s) - 1)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.s)-1 == this.ms.Top() {
		this.ms.Pop()
	}
	this.s = this.s[:len(this.s)-1]
}

func (this *MinStack) Top() int {
	return this.s[len(this.s)-1]
}

func (this *MinStack) GetMin() int {
	return this.s[this.ms.Top()]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

