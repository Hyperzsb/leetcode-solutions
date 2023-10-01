type Stack []int

func (s *Stack) Push(n int) {
    *s = append(*s, n)
}

func (s *Stack) Pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() int {
    return (*s)[len(*s)-1]
}

func find132pattern(nums []int) bool {
    stack := make(Stack, 0)
    third := math.MinInt32

    result := false
    for i := len(nums)-1; i >= 0; i-- {
        if nums[i] < third {
            result = true
            break
        }

        for len(stack) > 0 && stack.Top() < nums[i] {
            third = stack.Top()
            stack.Pop()
        }

        stack.Push(nums[i])
    }

    return result
}

