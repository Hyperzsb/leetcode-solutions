/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func construct(s, e int) []*TreeNode {
    result := make([]*TreeNode, 0)

    if s > e {
        return result
    }

    if s == e {
        result = append(result, &TreeNode{
            Val: s,
        })

        return result
    }

    for i := s; i <= e ; i++ {
        leftRoots := construct(s, i-1)
        rightRoots := construct(i+1, e)

        if len(leftRoots) == 0 && len(rightRoots) == 0 {
            result = append(result, &TreeNode{
                Val: i,
            })
        } else if len(leftRoots) == 0 && len(rightRoots) != 0 {
            for j := range rightRoots {
                result = append(result, &TreeNode{
                    Val: i,
                    Right: rightRoots[j],
                })
            }
        } else if len(leftRoots) != 0 && len(rightRoots) == 0 {
            for j := range leftRoots {
                result = append(result, &TreeNode{
                    Val: i,
                    Left: leftRoots[j],
                })
            }
        } else {
            for j := range leftRoots {
                for k := range rightRoots {
                    result = append(result, &TreeNode{
                        Val: i,
                        Left: leftRoots[j],
                        Right: rightRoots[k],
                    })
                }
            }
        }
    }

    return result
}

func generateTrees(n int) []*TreeNode {
    return construct(1, n)
}

