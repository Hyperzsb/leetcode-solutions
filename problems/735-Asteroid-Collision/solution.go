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

func asteroidCollision(asteroids []int) []int {
	stack := make(Stack, 0, len(asteroids))

	for _, val := range asteroids {
		if val > 0 {
			stack.Push(val)
		} else {
			for !stack.Empty() && stack.Top() > 0 && stack.Top() < -val {
				stack.Pop()
			}
			if stack.Empty() || stack.Top() < 0 {
				stack.Push(val)
			}
			if stack.Top() == -val {
				stack.Pop()
			}
		}
	}

	return stack
}

