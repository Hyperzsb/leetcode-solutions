/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue []*TreeNode

func (q *Queue) Push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() *TreeNode {
	return (*q)[0]
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	nodeQ := make(Queue, 0)

	result := make([][]int, 0)

	if root != nil {
		nodeQ.Push(root)
	}

	direction := false
	for len(nodeQ) > 0 {
		length := len(nodeQ)
		line := make([]int, length)

		for i := 0; i < length; i++ {
			node := nodeQ.Front()
			nodeQ.Pop()

			line[i] = node.Val

			if node.Left != nil {
				nodeQ.Push(node.Left)
			}

			if node.Right != nil {
				nodeQ.Push(node.Right)
			}
		}

		if direction {
			for i := 0; i < len(line)/2; i++ {
				line[i] ^= line[len(line)-1-i]
				line[len(line)-1-i] ^= line[i]
				line[i] ^= line[len(line)-1-i]
			}

			direction = false
		} else {
			direction = true
		}

		result = append(result, line)
	}

	return result
}

