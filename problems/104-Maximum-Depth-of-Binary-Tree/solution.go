/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, current int) int {
	if root == nil {
		return current - 1
	} else if root.Left == nil && root.Right == nil {
		return current
	} else {
		left := traverse(root.Left, current+1)
		right := traverse(root.Right, current+1)
		if left >= right {
			return left
		} else {
			return right
		}
	}
}

func maxDepth(root *TreeNode) int {
	return traverse(root, 1)
}

