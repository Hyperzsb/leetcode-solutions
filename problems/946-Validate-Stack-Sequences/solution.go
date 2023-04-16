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

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make(Stack, 0)
	pushedMap := make(map[int]bool)

	idxPush, idxPop := 0, 0
	for idxPop < len(popped) {
		if pushedMap[popped[idxPop]] {
			if stack.Top() == popped[idxPop] {
				stack.Pop()
				idxPop++
			} else {
				return false
			}
		} else {
			for idxPush < len(pushed) && pushed[idxPush] != popped[idxPop] {
				stack.Push(pushed[idxPush])
				pushedMap[pushed[idxPush]] = true
				idxPush++
			}

			stack.Push(pushed[idxPush])
			pushedMap[pushed[idxPush]] = true
			idxPush++
		}
	}

	return true
}

