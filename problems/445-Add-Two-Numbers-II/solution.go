/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
    prev, curr := (*ListNode)(nil), head

    for curr.Next != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    curr.Next = prev

    return curr
}

func length(head *ListNode) int {
    result := 0

    for head != nil {
        result++
        head = head.Next
    }

    return result
}

func add(head1, head2 *ListNode) *ListNode {
    if length(head1) < length(head2) {
        temp := head1
        head1 = head2
        head2 = temp
    }

    curr1, prev1, curr2 := head1, (*ListNode)(nil), head2

    sum, cum := 0, 0
    for curr1 != nil && curr2 != nil {
        sum = curr1.Val + curr2.Val + cum

        if sum >= 10 {
            cum = 1
            sum %= 10
        } else {
            cum = 0
        }

        curr1.Val = sum
        
        prev1 = curr1
        curr1 = curr1.Next
        curr2 = curr2.Next
    }

    for curr1 != nil {
        sum = curr1.Val + cum

        if sum >= 10 {
            cum = 1
            sum %= 10
        } else {
            cum = 0
        }

        curr1.Val = sum
        
        prev1 = curr1
        curr1 = curr1.Next
    }

    if cum == 1 {
        prev1.Next = &ListNode{
            Val: 1,
            Next: nil,
        }
    }

    return head1
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1, l2 = reverse(l1), reverse(l2)

    result := add(l1, l2)

    return reverse(result)
}

