type SmallestInfiniteSet struct {
	Queue []int
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{make([]int, 1)}
	set.Queue[0] = 1

	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.Queue) == 1 {
		this.Queue[0]++
		return this.Queue[0] - 1
	} else {
		result := this.Queue[len(this.Queue)-1]
		this.Queue = this.Queue[:len(this.Queue)-1]
		return result
	}
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.Queue[0] {
		return
	}

	for i := 1; i < len(this.Queue); i++ {
		if this.Queue[i] == num {
			return
		}
	}

	this.Queue = append(this.Queue, num)
	sort.Slice(this.Queue, func(a, b int) bool {
		return this.Queue[a] > this.Queue[b]
	})
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

