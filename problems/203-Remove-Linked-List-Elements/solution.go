/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	var pre, current, next *ListNode = nil, head, head.Next
	for current != nil {
		if current.Val == val {
			if pre == nil && next == nil {
				head = nil
				pre = nil
				current = current.Next
				next = nil
			} else if pre != nil && next == nil {
				pre.Next = nil
				current = current.Next
				next = nil
			} else if pre == nil && next != nil {
				head = current.Next
				pre = nil
				current = current.Next
				next = current.Next
			} else {
				pre.Next = next
				current = current.Next
				next = current.Next
			}
		} else {
			pre = current
			current = current.Next
			if current != nil {
				next = current.Next
			} else {
				next = nil
			}
		}
	}

	return head
}

