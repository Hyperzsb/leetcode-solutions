/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    newHead := &ListNode{
        Next: head,
    }

    prev, curr := newHead, head
    for i := 1; i < left; i++ {
        prev = curr
        curr = curr.Next
    }

    oldTail, newTail := prev, curr

    for i := left; i <= right; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    oldTail.Next = prev
    newTail.Next = curr

    return newHead.Next
}

