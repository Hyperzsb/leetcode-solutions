/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue []*TreeNode

func (q *Queue) Push(n *TreeNode) {
	*q = append(*q, n)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() *TreeNode {
	return (*q)[0]
}

func isCompleteTree(root *TreeNode) bool {
	queue := make(Queue, 0)
	queue.Push(root)

	level := 0
	maxNode := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for len(queue) > 0 {
		hasChildren := false
		lessThanTwoChildren := false
		length := len(queue)

		for i := 0; i < length; i++ {
			n := queue.Front()
			queue.Pop()

			if n.Left == nil && n.Right == nil {
				lessThanTwoChildren = true
			} else if n.Left == nil && n.Right != nil {
				return false
			} else if n.Left != nil && n.Right == nil {
				if lessThanTwoChildren {
					return false
				}

				hasChildren = true
				lessThanTwoChildren = true
				queue.Push(n.Left)
			} else {
				if lessThanTwoChildren {
					return false
				}

				hasChildren = true
				queue.Push(n.Left)
				queue.Push(n.Right)
			}
		}

		if hasChildren && length != maxNode[level] {
			return false
		}

		level++
	}

	return true
}

