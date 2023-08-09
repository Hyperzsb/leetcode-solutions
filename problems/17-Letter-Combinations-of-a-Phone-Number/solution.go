var digitToLetter = [][]byte{
    {},
    {},
    {'a', 'b', 'c'},
    {'d', 'e', 'f'},
    {'g', 'h', 'i'},
    {'j', 'k', 'l'},
    {'m', 'n', 'o'},
    {'p', 'q', 'r', 's'},
    {'t', 'u', 'v'},
    {'w', 'x', 'y', 'z'},
}

func backtraceLetter(curr []byte, digits []int, result *[][]byte) {
    if len(digits) == 0 {
        return
    }
    
    if len(curr) == len(digits)-1 {
        for _, letter := range digitToLetter[digits[len(curr)]] {
            newResult := make([]byte, len(curr))
            copy(newResult, curr)
            newResult = append(newResult, letter)
            *result = append(*result, newResult)
        }

        return
    }

    for _, letter := range digitToLetter[digits[len(curr)]] {
        curr = append(curr, letter)
        backtraceLetter(curr, digits, result)
        curr = curr[:len(curr)-1]
    }
}

// func backtraceDigit(curr []int, nums []int, result *[][]int) {
//     if len(nums) == 1 {
//         newResult := make([]int, len(curr))
//         copy(newResult, curr)
//         newResult = append(newResult, nums[0])
//         *result = append(*result, newResult)

//         return
//     }

//     for i := 0; i < len(nums); i++ {
//         curr = append(curr, nums[i])
        
//         newNums := make([]int, i)
//         copy(newNums, nums[:i])
//         newNums = append(newNums, nums[i+1:]...)

//         backtraceDigit(curr, newNums, result)

//         curr = curr[:len(curr)-1]
//     }
// }

func letterCombinations(digits string) []string {
    nums := make([]int, len(digits))
    for i := range digits {
        nums[i] = int(digits[i] - '0')
    }

    // digitCurr, digitResult := make([]int, 0), make([][]int, 0)
    // backtraceDigit(digitCurr, nums, &digitResult)

    // letterCurr, letterResult := make([]byte, 0), make([][]byte, 0)
    // for i := range digitResult {
    //     backtraceLetter(letterCurr, digitResult[i], &letterResult)
    // }

    letterCurr, letterResult := make([]byte, 0), make([][]byte, 0)
    backtraceLetter(letterCurr, nums, &letterResult)

    result := make([]string, len(letterResult))
    for i := range letterResult {
        result[i] = string(letterResult[i])
    }

    return result
}

