/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	current := head
	for current != nil {
		if current.Val > 100000 {
			return true
		}
		current.Val = 1000000
		current = current.Next
	}

	return false
}

