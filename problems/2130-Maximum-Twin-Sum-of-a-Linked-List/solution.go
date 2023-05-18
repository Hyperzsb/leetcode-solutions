/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func pairSum(head *ListNode) int {
	slow, fast := head, head

	for fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = slow.Next

	prev, curr := (*ListNode)(nil), slow
	for curr.Next != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	curr.Next = prev

	max := head.Val + curr.Val
	for curr != nil {
		if max < head.Val+curr.Val {
			max = head.Val + curr.Val
		}

		head = head.Next
		curr = curr.Next
	}

	return max
}

