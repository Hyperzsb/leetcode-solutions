/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root *TreeNode, trace *[]int, result *[]string) {
    *trace = append(*trace, root.Val)

    if root.Left == nil && root.Right == nil {
        str := fmt.Sprintf("%d", (*trace)[0])

        for i := 1; i < len(*trace); i++ {
            str += fmt.Sprintf("->%d", (*trace)[i])
        }

        *result = append(*result, str)
    }

    if root.Left != nil {
        traverse(root.Left, trace, result)
    }

    if root.Right != nil {
        traverse(root.Right, trace, result)
    }

    *trace = (*trace)[:len(*trace)-1]
}

func binaryTreePaths(root *TreeNode) []string {
    result := make([]string, 0)
    trace := make([]int, 0)
    traverse(root, &trace, &result)

    return result
}

