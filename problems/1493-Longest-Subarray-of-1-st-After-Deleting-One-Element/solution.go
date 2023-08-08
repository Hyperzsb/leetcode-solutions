func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func longestSubarray(nums []int) int {
    flag := true
    for i := range nums {
        if nums[i] == 1 {
            flag = false
            break
        }
    }

    if flag {
        return 0
    }

    nums = append(nums, 0)

    oneLists := make([][]int, 0)

    prevOne, oneCnt, zeroCnt := false, 0, 0
    for i := range nums {
        if nums[i] == 0 {
            if !prevOne {
                zeroCnt++
                continue
            }

            if zeroCnt > 1 {
                oneLists = append(oneLists, make([]int, 0))
                oneLists[len(oneLists)-1] = append(oneLists[len(oneLists)-1], oneCnt)
            } else {
                if len(oneLists) == 0 {
                    oneLists = append(oneLists, make([]int, 0))
                }
                oneLists[len(oneLists)-1] = append(oneLists[len(oneLists)-1], oneCnt)
            }

            prevOne = false
            zeroCnt = 1
        } else {
            if prevOne {
                oneCnt++
                continue
            }

            prevOne = true
            oneCnt = 1
        }
    }

    fmt.Println(oneLists)

    result := 0
    for i := range oneLists {
        if len(oneLists[i]) == 1 { 
            if len(oneLists) == 1 {
                if oneLists[i][0] == len(nums)-1 {
                    result = max(result, oneLists[i][0]-1)
                } else {
                    result = max(result, oneLists[i][0])
                }
            } else {
                result = max(result, oneLists[i][0])
            }

            continue
        }

        sum := oneLists[i][0]
        for j := 1; j < len(oneLists[i]); j++ {
            sum += oneLists[i][j]
            result = max(result, sum)
            sum -=oneLists[i][j-1]
        }
    }

    return result
}

