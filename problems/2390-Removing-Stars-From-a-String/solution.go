type Stack []byte

func (s *Stack) Push(char byte) {
	*s = append(*s, char)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func removeStars(s string) string {
	stack := make(Stack, 0)
	for _, char := range s {
		if char != '*' {
			stack.Push(byte(char))
		} else {
			stack.Pop()
		}
	}

	return string(stack)
}

