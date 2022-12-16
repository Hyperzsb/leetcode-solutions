type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	var numArray NumArray
	numArray.arr = make([]int, len(nums))
	copy(numArray.arr, nums)
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	result := 0
	for i := left; i <= right; i++ {
		result += this.arr[i]
	}
	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

