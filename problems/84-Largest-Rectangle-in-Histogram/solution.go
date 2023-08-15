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

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxArea(heights []int, s Stack) int {
    result := heights[s.Top()]

    for i := len(s)-1; i >= 1; i-- {
        result = max(result, (s.Top()-s[i-1])*heights[s[i]])
    }

    result = max(result, (s.Top()+1)*heights[s[0]])

    return result
}

func largestRectangleArea(heights []int) int {
    stack := make(Stack, 0)

    result := 0
    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] <= heights[stack.Top()] {
            stack.Pop()
        }

        stack.Push(i)

        r := maxArea(heights, stack)
        result = max(result, r)
    }

    return result
}

