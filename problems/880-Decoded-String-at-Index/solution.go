type Word struct {
    String string
    Length int64
    Factor int64
}

type Stack []*Word

func (s *Stack) Push(w *Word) {
    *s = append(*s, w)
}

func (s *Stack) Pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() *Word {
    return (*s)[len(*s)-1]
}

func decodeAtIndex(s string, k int) string {
    stack := make(Stack, 0)
    stack.Push(&Word{
        String: string(s[0]),
        Length: 1,
        Factor: 1,
    })

    for i := 1; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
            if stack.Top().Factor == 1 {
                stack.Top().String += string(s[i])
                stack.Top().Length += 1
            } else {
                stack.Push(&Word{
                    String: string(s[i]),
                    Length: stack.Top().Length*stack.Top().Factor+1,
                    Factor: 1,
                })
            }
        } else {
            stack.Top().Factor *= int64(s[i]-'0')
        }
    }

    result, t := "", int64(k)
    for len(stack) > 0 {
        t %= stack.Top().Length

        if t == 0 {
            t = stack.Top().Length
        }

        if t > stack.Top().Length-int64(len(stack.Top().String)) {
            idx := t-(stack.Top().Length-int64(len(stack.Top().String)))-1
            result = string(stack.Top().String[idx])
            break
        } else {
            stack.Pop()
        }
    }

    return result
}

