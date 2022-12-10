/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return 0 - num
	}
}

func calcSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	leftSum := calcSum(root.Left)
	rightSum := calcSum(root.Right)

	root.Val += leftSum + rightSum

	return root.Val
}

func findMin(root *TreeNode, target int, min *int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	if root.Left != nil {
		if abs(target-root.Left.Val) < abs(*min) {
			*min = target - root.Left.Val
		}

		findMin(root.Left, target, min)
	}

	if root.Right != nil {
		if abs(target-root.Right.Val) < abs(*min) {
			*min = target - root.Right.Val
		}

		findMin(root.Right, target, min)
	}
}

func maxProduct(root *TreeNode) int {
	sum := calcSum(root)

	min := sum / 2
	findMin(root, sum/2, &min)

	return ((sum/2 - min) * (sum - (sum/2 - min))) % 1000000007
}

