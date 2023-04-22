/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func traverse(root *TreeNode, direction int, lengthMap map[*TreeNode]map[int]int) int {
	if root == nil {
		return 0
	}

	if _, ok := lengthMap[root]; !ok {
		lengthMap[root] = make(map[int]int)
	}

	result := 0
	if direction == -1 {
		if val, ok := lengthMap[root][-1]; ok {
			result = val
		} else {
			result = traverse(root.Right, 1, lengthMap) + 1
			lengthMap[root][-1] = result
		}
	}

	if direction == 1 {
		if val, ok := lengthMap[root][1]; ok {
			result = val
		} else {
			result = traverse(root.Left, -1, lengthMap) + 1
			lengthMap[root][1] = result
		}
	}

	if direction == 0 {
		if val, ok := lengthMap[root][0]; ok {
			result = val
		} else {
			result = max(result, traverse(root.Left, 0, lengthMap))
			result = max(result, traverse(root.Right, 0, lengthMap))
			result = max(result, traverse(root.Left, -1, lengthMap))
			result = max(result, traverse(root.Right, 1, lengthMap))
			lengthMap[root][0] = result
		}
	}

	return result
}

func longestZigZag(root *TreeNode) int {
	lengthMap := make(map[*TreeNode]map[int]int)
	return traverse(root, 0, lengthMap)
}

