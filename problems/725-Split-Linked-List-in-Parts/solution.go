/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func lengthOf(head *ListNode) int {
    result := 0
    for head != nil {
        result++
        head = head.Next
    }

    return result
}

func split(head *ListNode, l int) (*ListNode, *ListNode) {
    if head == nil {
        return nil, nil
    }

    if l < 1 {
        return nil, head
    }

    prev, curr := (*ListNode)(nil), head
    for i := 0; i < l; i++ {
        prev = curr
        curr = curr.Next
    }

    prev.Next = nil

    return head, curr
}

func splitListToParts(head *ListNode, k int) []*ListNode {
    length := lengthOf(head)

    result := make([]*ListNode, 0)
    for i := 0; i < length%k; i++ {
        p, h := split(head, length/k+1)
        result = append(result, p)
        head = h
    }

    for i := 0; i < k-length%k; i++ {
        p, h := split(head, length/k)
        result = append(result, p)
        head = h
    }

    return result
}

