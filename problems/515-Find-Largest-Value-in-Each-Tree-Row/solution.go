/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue []*TreeNode

func (q *Queue) Push(tn *TreeNode) {
    *q = append(*q, tn)
}

func (q *Queue) Pop() {
    *q = (*q)[1:]
}

func (q *Queue) Front() *TreeNode {
    return (*q)[0]
}

func largestValues(root *TreeNode) []int {
    result := make([]int, 0)

    if root == nil {
        return result
    }

    q := make(Queue, 0)
    q.Push(root)
    
    for len(q) > 0 {
        length := len(q)

        max := math.MinInt32
        for i := 0; i < length; i++ {
            tn := q.Front()
            q.Pop()

            if max < tn.Val {
                max = tn.Val
            }

            if tn.Left != nil {
                q.Push(tn.Left)
            }

            if tn.Right != nil {
                q.Push(tn.Right)
            }
        }

        result = append(result, max)
    }

    return result
}

