type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Reverse() {
	for i := 0; i < len(*s)/2; i++ {
		(*s)[i] ^= (*s)[len(*s)-1-i]
		(*s)[len(*s)-1-i] ^= (*s)[i]
		(*s)[i] ^= (*s)[len(*s)-1-i]
	}
}

func NextNum(s *string, idx int) (int, int) {
	for (*s)[idx] == ' ' {
		idx++
	}

	num := 0
	for idx < len(*s) && (*s)[idx] >= '0' && (*s)[idx] <= '9' {
		num += int((*s)[idx] - '0')
		num *= 10
		idx++
	}

	for idx < len(*s) && (*s)[idx] == ' ' {
		idx++
	}

	return num / 10, idx
}

func calculate(s string) int {
	numStack, opStack := make(Stack, 0), make(Stack, 0)

	for num, i := 0, 0; i < len(s); i++ {
		num, i = NextNum(&s, i)

		if !opStack.Empty() && (opStack.Top() == 3 || opStack.Top() == 4) {
			if opStack.Top() == 3 {
				num = numStack.Top() * num
			} else {
				num = numStack.Top() / num
			}
			opStack.Pop()
			numStack.Pop()
			numStack.Push(num)
		} else {
			numStack.Push(num)
		}

		if i == len(s) {
			break
		}

		switch s[i] {
		case '+':
			opStack.Push(1)
		case '-':
			opStack.Push(2)
		case '*':
			opStack.Push(3)
		case '/':
			opStack.Push(4)
		}
	}

	numStack.Reverse()
	opStack.Reverse()

	for !opStack.Empty() {
		a := numStack.Top()
		numStack.Pop()
		b := numStack.Top()
		numStack.Pop()

		if opStack.Top() == 1 {
			numStack.Push(a + b)
		} else {
			numStack.Push(a - b)
		}

		opStack.Pop()
	}

	return numStack.Top()
}

