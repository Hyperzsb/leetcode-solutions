/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := TreeNode{Val: postorder[len(postorder)-1]}

	var leftInorder, rightInorder []int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			leftInorder = inorder[:i]
			rightInorder = inorder[i+1:]
		}
	}

	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(leftInorder)+len(rightInorder)]

	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)

	return &root
}

