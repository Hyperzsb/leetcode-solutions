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

func dfs(root, target *TreeNode, stack *[]*TreeNode, status *bool) {
    if root == nil {
        return
    }

    if *status {
        return
    }

    *stack = append(*stack, root)

    if root == target {
        *status = true
        return
    }

    dfs(root.Left, target, stack, status)
    dfs(root.Right, target, stack, status)

    if !(*status) {
        *stack = (*stack)[:len(*stack)-1]
    }
}

func bfs(root *TreeNode, k int, maskMap map[int]bool) []int {
    q := make(Queue, 0)
    q.Push(root.Left)
    q.Push(root.Right)

    depth := 1
    for len(q) > 0 {
        nodes := make([]int, 0)

        length := len(q)
        for i := 0; i < length; i++ {
            tn := q.Front()
            q.Pop()

            if tn == nil || maskMap[tn.Val] {
                continue
            }

            nodes = append(nodes, tn.Val)

            q.Push(tn.Left)
            q.Push(tn.Right)
        }

        if depth == k {
            return nodes
        }

        depth++
    }

    return nil
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    stack, status := make([]*TreeNode, 0), false
    dfs(root, target, &stack, &status)

    maskMap := make(map[int]bool)

    for i := range stack {
        if len(stack)-1-i >= k { 
            continue
        }

        maskMap[stack[i].Val] = true
    }

    result := make([]int, 0)
    for i := range stack {
        if len(stack)-1-i > k { 
            continue
        }

        if len(stack)-1-i == k {
            result = append(result, stack[i].Val)
        }

        nodes := bfs(stack[i], k-(len(stack)-1-i), maskMap)
        for j := range nodes {
            result = append(result, nodes[j])
        }
    }

    return result
}

