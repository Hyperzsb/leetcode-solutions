/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil || root.Right != nil {
		traverse(root.Left, leaves)
		traverse(root.Right, leaves)
	} else {
		*leaves = append(*leaves, root.Val)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := make([]int, 0), make([]int, 0)
	traverse(root1, &leaves1)
	traverse(root2, &leaves2)

	len1, len2 := len(leaves1), len(leaves2)
	if len1 == len2 {
		for i := 0; i < len1; i++ {
			if leaves1[i] != leaves2[i] {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

