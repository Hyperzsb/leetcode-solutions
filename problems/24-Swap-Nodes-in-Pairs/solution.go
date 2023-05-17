/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast, newHead, prev := head, head.Next, head.Next, (*ListNode)(nil)

	for {
		slow.Next = fast.Next
		fast.Next = slow

		if prev != nil {
			prev.Next = fast
		}
		prev = slow

		if slow.Next == nil || slow.Next.Next == nil {
			break
		}

		slow = slow.Next
		fast = slow.Next
	}

	return newHead
}

