/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func traverse(root *TreeNode, pre, res *int) {
	if root == nil {
		return
	}

	traverse(root.Left, pre, res)

	if diff(root.Val, *pre) < *res {
		*res = diff(root.Val, *pre)
	}

	*pre = root.Val

	traverse(root.Right, pre, res)
}

func getMinimumDifference(root *TreeNode) int {
	pre, res := 1000000, 1000000
	traverse(root, &pre, &res)

	return res
}

