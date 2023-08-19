/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := &ListNode{
        Next: head,
    }

    lessHead := newHead
    for lessHead.Next != nil {
        if lessHead.Next.Val >= x {
            break
        }

        lessHead = lessHead.Next
    }

    prev, curr := lessHead, lessHead.Next
    for curr != nil {
        if curr.Val < x {
            prev.Next = curr.Next

            curr.Next = lessHead.Next
            lessHead.Next = curr
            lessHead = curr

            curr = prev.Next
        } else {
            prev = curr
            curr = curr.Next
        }
    }

    return newHead.Next
}

