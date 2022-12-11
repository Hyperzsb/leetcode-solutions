/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, flag bool, max *int) int {
	if flag {
		if root.Left == nil && root.Right == nil {
			if root.Val > *max {
				*max = root.Val
			}
			return root.Val
		} else if root.Left != nil && root.Right == nil {
			left := traverse(root.Left, true, max)
			m := 0
			if left <= 0 {
				m = root.Val
			} else {
				m = root.Val + left
			}
			if m > *max {
				*max = m
			}
			return m
		} else if root.Left == nil && root.Right != nil {
			right := traverse(root.Right, true, max)
			m := 0
			if right <= 0 {
				m = root.Val
			} else {
				m = root.Val + right
			}
			if m > *max {
				*max = m
			}
			return m
		} else {
			left := traverse(root.Left, true, max)
			right := traverse(root.Right, true, max)
			m := 0
			if left > right {
				if left <= 0 {
					m = root.Val
				} else {
					m = root.Val + left
				}
				if m > *max {
					*max = m
				}
				return m
			} else {
				if right <= 0 {
					m = root.Val
				} else {
					m = root.Val + right
				}
				if m > *max {
					*max = m
				}
				return m
			}
		}
	} else {
		if root.Left == nil && root.Right == nil {
			if root.Val > *max {
				*max = root.Val
			}
			return 0
		} else if root.Left != nil && root.Right == nil {
			left := traverse(root.Left, true, max)
			m := 0
			if left <= 0 {
				m = root.Val
			} else {
				m = root.Val + left
			}
			if m > *max {
				*max = m
			}
		} else if root.Left == nil && root.Right != nil {
			right := traverse(root.Right, true, max)
			m := 0
			if right <= 0 {
				m = root.Val
			} else {
				m = root.Val + right
			}
			if m > *max {
				*max = m
			}
		} else {
			left := traverse(root.Left, true, max)
			right := traverse(root.Right, true, max)
			m := 0
			if left > right {
				if left <= 0 {
					m = root.Val
				} else {
					m = root.Val + left
				}
				if m > *max {
					*max = m
				}
			} else {
				if right <= 0 {
					m = root.Val
				} else {
					m = root.Val + right
				}
				if m > *max {
					*max = m
				}
			}
			if left+root.Val+right > m {
				if left+root.Val+right > *max {
					*max = left + root.Val + right
				}
			}
		}
		if root.Left != nil {
			traverse(root.Left, false, max)
		}
		if root.Right != nil {
			traverse(root.Right, false, max)
		}
		return 0
	}
}

func maxPathSum(root *TreeNode) int {
	max := -2100000000
	traverse(root, false, &max)
	return max
}

