/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func length(head *ListNode) int {
    l := 0
    for head != nil {
        head = head.Next
        l++
    }

    return l
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
    prev, curr := (*ListNode)(nil), head

    for i := 0; i < k-1; i++ {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    next := curr.Next
    curr.Next = prev
    head.Next = next

    return curr, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    l := length(head)
    if l < k {
        return head
    }

    head, tail := reverse(head, k)

    for i := 2; i*k <= l; i++ {
        h, t := reverse(tail.Next, k)
        tail.Next = h
        tail = t
    }

    return head
}

