/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}

		if isSameTree(p.Left, q.Left) == false {
			return false
		} else {
			return isSameTree(p.Right, q.Right)
		}
	} else if p == nil && q == nil {
		return true
	} else {
		return false
	}
}

