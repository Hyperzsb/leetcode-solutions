/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, parent **TreeNode, count, max *int, result *map[int]int) {
	if root == nil {
		return
	}

	traverse(root.Left, parent, count, max, result)

	if *parent == nil {
		*count = 1
	} else if root.Val == (*parent).Val {
		*count += 1
	} else {
		*count = 1
	}

	*parent = root

	if *count == *max {
		(*result)[root.Val] = *count
	}

	if *count > *max {
		*max = *count

		for key, _ := range *result {
			delete(*result, key)
		}

		(*result)[root.Val] = *count
	}

	traverse(root.Right, parent, count, max, result)
}

func findMode(root *TreeNode) []int {
	vals := make(map[int]int)

	count, max := 0, 0
	var parent *TreeNode
	parent = nil

	traverse(root, &parent, &count, &max, &vals)

	result := make([]int, 0, len(vals))
	for key, _ := range vals {
		result = append(result, key)
	}

	return result
}

