/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, cum := 0, 0
	previous, current1, current2 := l1, l1, l2

	for current1 != nil && current2 != nil {
		sum = (current1.Val + current2.Val + cum) % 10
		cum = (current1.Val + current2.Val + cum) / 10
		current1.Val = sum
		previous = current1
		current1 = current1.Next
		current2 = current2.Next
	}

	for current2 != nil {
		previous.Next = current2
		sum = (current2.Val + cum) % 10
		cum = (current2.Val + cum) / 10
		current2.Val = sum
		previous = current2
		current2 = current2.Next
	}

	for current1 != nil {
		previous.Next = current1
		sum = (current1.Val + cum) % 10
		cum = (current1.Val + cum) / 10
		current1.Val = sum
		previous = current1
		current1 = current1.Next
	}

	if cum > 0 {
		node := ListNode{cum, nil}
		previous.Next = &node
	}

	return l1
}

