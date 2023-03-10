/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	slowPtr, fastPtr := head, head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if slowPtr == fastPtr {
			headPtr, collisionPtr := head, slowPtr
			for headPtr != collisionPtr {
				headPtr = headPtr.Next
				collisionPtr = collisionPtr.Next
			}

			return collisionPtr
		}
	}

	return nil
}

