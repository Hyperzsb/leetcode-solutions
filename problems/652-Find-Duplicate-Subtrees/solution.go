/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func stringify(nums []int) string {
	var result string
	for _, num := range nums {
		result = fmt.Sprintf("%v%v,", result, num)
	}

	return result
}

func preorderTraverse(root *TreeNode, direction int, preorder *[]int, preorderMap *map[*TreeNode]string) {
	if root == nil {
		return
	}

	length := len(*preorder)

	*preorder = append(*preorder, root.Val+200*direction)
	preorderTraverse(root.Left, -1, preorder, preorderMap)
	preorderTraverse(root.Right, 1, preorder, preorderMap)

	(*preorder)[length] -= 200 * direction
	thisPreorder := make([]int, len(*preorder)-length)
	copy(thisPreorder, (*preorder)[length:])
	(*preorder)[length] += 200 * direction

	(*preorderMap)[root] = stringify(thisPreorder)
}

func inorderTraverse(root *TreeNode, direction int, inorder *[]int, inorderMap *map[*TreeNode]string) {
	if root == nil {
		return
	}

	length := len(*inorder)

	inorderTraverse(root.Left, -1, inorder, inorderMap)
	idx := len(*inorder)
	*inorder = append(*inorder, root.Val+200*direction)
	inorderTraverse(root.Right, 1, inorder, inorderMap)

	(*inorder)[idx] -= 200 * direction
	thisInorder := make([]int, len(*inorder)-length)
	copy(thisInorder, (*inorder)[length:])
	(*inorder)[idx] += 200 * direction

	(*inorderMap)[root] = stringify(thisInorder)
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	preorder := make([]int, 0)
	preorderMap := make(map[*TreeNode]string)
	preorderTraverse(root, 0, &preorder, &preorderMap)

	inorder := make([]int, 0)
	inorderMap := make(map[*TreeNode]string)
	inorderTraverse(root, 0, &inorder, &inorderMap)

	subtrees := make(map[string]int)
	result := make([]*TreeNode, 0)
	for node, pre := range preorderMap {
		in := inorderMap[node]
		key := fmt.Sprintf("%v|%v", pre, in)

		if subtrees[key] == 0 {
			subtrees[key] = 1
		} else if subtrees[key] == 1 {
			result = append(result, node)
			subtrees[key] = 2
		} else {
			continue
		}
	}

	return result
}

