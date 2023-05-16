/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	fast, slow := head, head

	for i := 1; i < k; i++ {
		fast = fast.Next
	}

	firstK := fast

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	lastK := slow

	temp := firstK.Val
	firstK.Val = lastK.Val
	lastK.Val = temp

	return head
}

