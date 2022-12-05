type Stack struct {
	stack []int
}

func (stack *Stack) Size() int {
	return len(stack.stack)
}

func (stack *Stack) Push(num int) {
	stack.stack = append(stack.stack, num)
}

func (stack *Stack) Pop() int {
	top := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return top
}

func (stack *Stack) Top() int {
	return stack.stack[len(stack.stack)-1]
}

func isValid(s string) bool {
	stack := Stack{make([]int, 0)}

	length := len(s)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(':
			stack.Push(1)

		case '[':
			stack.Push(2)

		case '{':
			stack.Push(3)

		case ')':
			if stack.Size() == 0 || stack.Top() != 1 {
				return false
			} else {
				stack.Pop()
			}

		case ']':
			if stack.Size() == 0 || stack.Top() != 2 {
				return false
			} else {
				stack.Pop()
			}

		case '}':
			if stack.Size() == 0 || stack.Top() != 3 {
				return false
			} else {
				stack.Pop()
			}
		}
	}

	if stack.Size() > 0 {
		return false
	} else {
		return true
	}

}

