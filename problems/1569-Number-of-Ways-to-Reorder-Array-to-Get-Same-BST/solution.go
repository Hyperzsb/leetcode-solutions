type MTreeNode struct {
    Val int
    NCnt int
    PCnt int
    Left *MTreeNode
    Right *MTreeNode
}

func construct(nums []int) *MTreeNode {
    root := &MTreeNode{
        Val: nums[0],
    }

    for i := 1; i < len(nums); i++ {
        curr := root
        for {
            if nums[i] > curr.Val {
                if curr.Right == nil {
                    curr.Right = &MTreeNode{
                        Val: nums[i],
                    }

                    break
                } else {
                    curr = curr.Right
                }

                continue
            }

            if nums[i] < curr.Val {
                if curr.Left == nil {
                    curr.Left = &MTreeNode{
                        Val: nums[i],
                    }

                    break
                } else {
                    curr = curr.Left
                }

                continue
            }
        }
    }

    return root
}

func traverse(root *MTreeNode, dp [][]int) {
    if root.Left == nil && root.Right == nil {
        root.NCnt = 1
        root.PCnt = 1
    } else if root.Left == nil && root.Right != nil {
        traverse(root.Right, dp)

        root.NCnt = root.Right.NCnt + 1
        root.PCnt = root.Right.PCnt
    } else if root.Left != nil && root.Right == nil {
        traverse(root.Left, dp)

        root.NCnt = root.Left.NCnt + 1
        root.PCnt = root.Left.PCnt
    } else {
        traverse(root.Left, dp)
        traverse(root.Right, dp)

        root.NCnt = root.Left.NCnt + root.Right.NCnt + 1
        root.PCnt = ((root.Left.PCnt * root.Right.PCnt) % (1e9 + 7) * DP(root.Left.NCnt, root.Right.NCnt, dp)) % (1e9 + 7)
    }

    // fmt.Println(root.Val, root.NCnt, root.PCnt)
}

func DP(a, b int, dp [][]int) int {


    if a == 0 {
        return 1
    }

    if dp[a][b] != 0 {
        return dp[a][b]
    }

    val := 0
    for i := b; i >= 0; i-- {
        if dp[a][i] != 0 {
            val = (val + dp[a][i]) % (1e9 + 7)
            break
        } else if dp[a-1][i] != 0 {
            val = (val + dp[a-1][i]) % (1e9 + 7)
        } else {
            val = (val + DP(a-1, i, dp)) % (1e9 + 7)
        }
    }   

    dp[a][b] = val
    
    return val
}

func numOfWays(nums []int) int {
    root := construct(nums)

    dp := make([][]int, len(nums)+1)
    for i := range dp {
        dp[i] = make([]int, len(nums)+1)
    }

    traverse(root, dp)

    return root.PCnt - 1
}

