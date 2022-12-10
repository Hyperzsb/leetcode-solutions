/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, max, min *int) {
	if root == nil {
		return
	}

	if root.Val > *max {
		*max = root.Val
	}

	if root.Val < *min {
		*min = root.Val
	}

	leftMax, leftMin := *max, *min
	traverse(root.Left, &leftMax, &leftMin)

	rightMax, rightMin := *max, *min
	traverse(root.Right, &rightMax, &rightMin)

	if leftMax-leftMin >= rightMax-rightMin {
		*max = leftMax
		*min = leftMin
	} else {
		*max = rightMax
		*min = rightMin
	}
}

func maxAncestorDiff(root *TreeNode) int {
	max, min := -1, 1000000
	traverse(root, &max, &min)
	return max - min
}

