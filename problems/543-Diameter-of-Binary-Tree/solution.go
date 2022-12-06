/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode, current int, diameter *int) int {
	if root.Left == nil && root.Right == nil {
		return current
	} else if root.Left != nil && root.Right == nil {
		leftHeight := height(root.Left, current+1, diameter)
		if leftHeight-current > *diameter {
			*diameter = leftHeight - current
		}
		return leftHeight
	} else if root.Left == nil && root.Right != nil {
		rightHeight := height(root.Right, current+1, diameter)
		if rightHeight-current > *diameter {
			*diameter = rightHeight - current
		}
		return rightHeight
	} else {
		leftHeight := height(root.Left, current+1, diameter)
		rightHeight := height(root.Right, current+1, diameter)
		if leftHeight-current+rightHeight-current > *diameter {
			*diameter = leftHeight - current + rightHeight - current
		}
		if leftHeight >= rightHeight {
			return leftHeight
		} else {
			return rightHeight
		}
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	height(root, 0, &diameter)

	return diameter
}

