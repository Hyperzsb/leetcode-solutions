/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, previous, min *int) {
	if root == nil {
		return
	}

	traverse(root.Left, previous, min)

	if root.Val-*previous < *min {
		*min = root.Val - *previous
	}

	*previous = root.Val

	traverse(root.Right, previous, min)
}

func minDiffInBST(root *TreeNode) int {
	previous, min := -100000, 100000
	traverse(root, &previous, &min)

	return min
}

