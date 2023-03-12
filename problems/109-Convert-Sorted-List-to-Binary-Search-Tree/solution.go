/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	previous, current := (*ListNode)(nil), head
	for i := 0; i < length/2; i++ {
		previous = current
		current = current.Next
	}

	root := TreeNode{Val: current.Val}

	if previous != nil {
		previous.Next = nil
		root.Left = sortedListToBST(head)
	} else {
		root.Left = nil
	}
	root.Right = sortedListToBST(current.Next)

	return &root
}

