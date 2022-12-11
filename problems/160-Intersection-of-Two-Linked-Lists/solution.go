/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	currentA, currentB := headA, headB

	var result *ListNode = nil

	for currentA != nil {
		currentA.Val += 1000000
		currentA = currentA.Next
	}

	for currentB != nil {
		if currentB.Val > 100000 {
			result = currentB
			break
		}
		currentB = currentB.Next
	}

	currentA = headA
	for currentA != nil {
		currentA.Val -= 1000000
		currentA = currentA.Next
	}

	return result
}

