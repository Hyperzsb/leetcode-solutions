/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, currentSum int, targetSum int, flag bool) int {
	if root == nil {
		return 0
	}

	sum := 0

	if flag == false {
		if root.Val == targetSum {
			sum++
		}
		sum += traverse(root.Left, targetSum-root.Val, targetSum, true)
		sum += traverse(root.Right, targetSum-root.Val, targetSum, true)
		sum += traverse(root.Left, targetSum, targetSum, false)
		sum += traverse(root.Right, targetSum, targetSum, false)
	} else {
		if root.Val == currentSum {
			sum++
		}
		sum += traverse(root.Left, currentSum-root.Val, targetSum, true)
		sum += traverse(root.Right, currentSum-root.Val, targetSum, true)
	}

	return sum
}

func pathSum(root *TreeNode, targetSum int) int {
	return traverse(root, targetSum, targetSum, false)
}

