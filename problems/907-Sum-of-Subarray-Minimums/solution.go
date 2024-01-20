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

func sumSubarrayMins(arr []int) int {
    leftDist, rightDist := make([]int, len(arr)), make([]int, len(arr))
    leftStack, rightStack := make(Stack, 0), make(Stack, 0)
    for i := range arr {
        for len(leftStack) > 0 && arr[leftStack.Top()] > arr[i] {
            leftStack.Pop()
        }

        if len(leftStack) > 0 {
            leftDist[i] = leftStack.Top()
        } else {
            leftDist[i] = -1
        }

        leftStack.Push(i)

        for len(rightStack) > 0 && arr[rightStack.Top()] >= arr[len(arr)-1-i] {
            rightStack.Pop()
        }

        if len(rightStack) > 0 {
            rightDist[len(arr)-1-i] = rightStack.Top()
        } else {
            rightDist[len(arr)-1-i] = len(arr)
        }

        rightStack.Push(len(arr)-1-i)
    }

    result := 0
    for i := range arr {
        result += (i - leftDist[i]) * (rightDist[i] - i) % (1e9 + 7) * arr[i] % (1e9 + 7)
        result %= 1e9 + 7
    }

    return result
}

