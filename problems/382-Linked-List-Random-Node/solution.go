/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Head *ListNode
	Size int
}

func Constructor(head *ListNode) Solution {
	size := 0
	current := head
	for current != nil {
		size++
		current = current.Next
	}

	return Solution{Head: head, Size: size}
}

func (this *Solution) GetRandom() int {
	length := rand.Intn(this.Size)
	current := this.Head
	for i := 1; i <= length; i++ {
		current = current.Next
	}

	return current.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

