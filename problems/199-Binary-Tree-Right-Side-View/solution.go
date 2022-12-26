/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Node struct {
	Node  *TreeNode
	Depth int
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	nodesAvail := make([]Node, 1)
	nodesAvail[0] = Node{root, 0}
	currentDepth := 0
	for i := 0; i < len(nodesAvail); i++ {
		if nodesAvail[i].Depth > currentDepth {
			result = append(result, nodesAvail[i-1].Node.Val)
			currentDepth = nodesAvail[i].Depth
		}
		if nodesAvail[i].Node.Left != nil {
			nodesAvail = append(nodesAvail, Node{nodesAvail[i].Node.Left, nodesAvail[i].Depth + 1})
		}
		if nodesAvail[i].Node.Right != nil {
			nodesAvail = append(nodesAvail, Node{nodesAvail[i].Node.Right, nodesAvail[i].Depth + 1})
		}
	}
	result = append(result, nodesAvail[len(nodesAvail)-1].Node.Val)

	return result
}

