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

func (q *Queue) Front() *TreeNode {
	return (*q)[0]
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func maxLevelSum(root *TreeNode) int {
	q := make(Queue, 0)
	q.Push(root)

	max, level, result := int(-1e9), 1, 1
	for len(q) > 0 {
		length, sum := len(q), 0
		for i := 0; i < length; i++ {
			curr := q.Front()
			q.Pop()

			sum += curr.Val

			if curr.Left != nil {
				q.Push(curr.Left)
			}

			if curr.Right != nil {
				q.Push(curr.Right)
			}
		}

		if sum > max {
			max = sum
			result = level
		}

		level++
	}

	return result
}

