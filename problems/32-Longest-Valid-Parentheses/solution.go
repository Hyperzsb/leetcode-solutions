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

func longestValidParentheses(s string) int {
    result := 0

    if len(s) < 2 {
        return 0
    }

    stack := make(Stack, 0)
    stack.Push(0)

    for i := 1; i < len(s); i++ {
        if s[i] == '(' {
            stack.Push(i)
            continue
        } 

        if len(stack) == 0 || s[stack.Top()] == ')' {
            stack.Push(i)
            continue
        }

        stack.Pop()

        if len(stack) == 0 {
            result = max(result, i+1)
        } else {
            result = max(result, i-stack.Top())
        }
    }

    return result
}

