/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, current *int, target int, result *int) {
	if root == nil || *current == target {
		return
	}

	if root.Left == nil && root.Right == nil {
		*current++
		if *current == target {
			*result = root.Val
		}
		return
	} else {
		traverse(root.Left, current, target, result)
		if *current == target {
			return
		}
		*current++
		if *current == target {
			*result = root.Val
			return
		}
		traverse(root.Right, current, target, result)
	}
}

func kthSmallest(root *TreeNode, k int) int {
	init, result := 0, 0
	traverse(root, &init, k, &result)

	return result
}

