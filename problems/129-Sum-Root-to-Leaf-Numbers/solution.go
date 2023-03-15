/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, current, result *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*current = *current*10 + root.Val
		*result += *current
		*current = (*current - root.Val) / 10
		return
	}

	*current = *current*10 + root.Val
	traverse(root.Left, current, result)
	traverse(root.Right, current, result)
	*current = (*current - root.Val) / 10
}

func sumNumbers(root *TreeNode) int {
	current, result := 0, 0
	traverse(root, &current, &result)

	return result
}

