/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func allPossibleFBT(n int) []*TreeNode {
    result := make([]*TreeNode, 0)

    if n == 1 {
        result = append(result, &TreeNode{
            Val: 0,
            Left: nil,
            Right: nil,
        })

        return result
    }

    for i := 1; i <= n-2; i += 2 {
        leftRoots := allPossibleFBT(i)
        rightRoots := allPossibleFBT(n-1-i)

        for j := range leftRoots {
            for k := range rightRoots {
                result = append(result, &TreeNode{
                    Val: 0,
                    Left: leftRoots[j],
                    Right: rightRoots[k],
                })
            }
        }
    }

    return result
}

