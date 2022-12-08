/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		if check(left.Left, right.Right) == false {
			return false
		}
		if check(left.Right, right.Left) == false {
			return false
		}
		return true
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

