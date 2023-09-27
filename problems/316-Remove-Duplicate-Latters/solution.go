type Stack []byte

func (s *Stack) Push(c byte) {
    *s = append(*s, c)
}

func (s *Stack) Pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() byte {
    return (*s)[len(*s)-1]
}

func removeDuplicateLetters(s string) string {
    lastIdx := make(map[byte]int)
    for i := range s {
        lastIdx[s[i]] = i
    }

    added, stack := make(map[byte]bool), make(Stack, 0)
    for i := range s {
        if added[s[i]] {
            continue
        }

        for len(stack) > 0 && stack.Top() > s[i] && i < lastIdx[stack.Top()] {
            added[stack.Top()] = false
            stack.Pop()
        }

        stack.Push(s[i])
        added[s[i]] = true
    }

    return string(stack)
}

