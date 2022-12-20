/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, isLeft bool, sum *int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if isLeft {
			*sum += root.Val
		}
		return
	}

	traverse(root.Left, true, sum)
	traverse(root.Right, false, sum)
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	traverse(root, false, &sum)
	return sum
}

