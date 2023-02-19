/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := root.Left, root.Right

	invertTree(left)
	invertTree(right)

	root.Left = right
	root.Right = left

	return root
}

