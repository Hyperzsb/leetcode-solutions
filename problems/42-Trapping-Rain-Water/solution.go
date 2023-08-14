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

func trap(height []int) int {
    if len(height) < 3 {
        return 0
    }

    stack := make(Stack, 0)
    stack.Push(0)

    result := 0
    for i := 1; i < len(height); i++ {
        for len(stack) > 0 && height[i] > height[stack.Top()] {
            curr := stack.Top()
            stack.Pop()

            if (len(stack) > 0) {
                result += (min(height[i], height[stack.Top()]) - height[curr]) * (i - stack.Top() - 1)
            }
        }

        stack.Push(i)
    }

    return result
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

