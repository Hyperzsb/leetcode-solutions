type Number struct {
    Cnt int
    Children [][]int
}

func traverse(n int, nums map[int]*Number) int {
    num := nums[n]
    if num.Cnt != -1 {
        return num.Cnt
    }

    num.Cnt = 1

    for i := range num.Children {
        leftCnt := traverse(num.Children[i][0], nums)
        rightCnt := traverse(num.Children[i][1], nums)
        if num.Children[i][0] == num.Children[i][1] {
            num.Cnt = (num.Cnt + leftCnt * rightCnt) % (1e9 + 7)
        } else {
            num.Cnt = (num.Cnt + leftCnt * rightCnt * 2) % (1e9 + 7)
        }
    }

    return num.Cnt
}

func numFactoredBinaryTrees(arr []int) int {
    sort.Slice(arr, func(a, b int) bool {
        return arr[a] < arr[b]
    })

    nums := make(map[int]*Number)
    for i := range arr {
        nums[arr[i]] = &Number{
            Cnt: -1,
            Children: make([][]int, 0),
        }
    }

    for i := 0; i < len(arr); i++ {
        for j := i; j < len(arr); j++ {
            if num, ok := nums[arr[i]*arr[j]]; ok {
                num.Children = append(num.Children, []int{arr[i], arr[j]})
            }
        }
    }

    result := 0
    for i := range arr {
        result = (result + traverse(arr[i], nums)) % (1e9 + 7)
    }

    return result
}

