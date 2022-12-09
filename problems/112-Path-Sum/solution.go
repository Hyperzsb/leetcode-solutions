/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, current, target int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if current+root.Val == target {
			return true
		} else {
			return false
		}
	} else {
		if traverse(root.Left, current+root.Val, target) {
			return true
		} else {
			return traverse(root.Right, current+root.Val, target)
		}
	}
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return traverse(root, 0, targetSum)
}

