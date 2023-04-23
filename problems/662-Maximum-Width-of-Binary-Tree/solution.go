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

func widthOfBinaryTree(root *TreeNode) int {
	queue := make(Queue, 0)
	root.Val = 1
	queue.Push(root)

	result := 0
	for len(queue) > 0 {
		length := len(queue)

		firstNode, lastNode := (*TreeNode)(nil), (*TreeNode)(nil)
		for i := 0; i < length; i++ {
			curr := queue.Front()
			queue.Pop()

			if firstNode == nil {
				firstNode = curr
				lastNode = curr
			} else {
				lastNode = curr
			}

			if curr.Left != nil {
				curr.Left.Val = curr.Val * 2
				queue.Push(curr.Left)
			}

			if curr.Right != nil {
				curr.Right.Val = curr.Val*2 + 1
				queue.Push(curr.Right)
			}
		}

		if lastNode.Val-firstNode.Val+1 > result {
			result = lastNode.Val - firstNode.Val + 1
		}
	}

	return result
}

