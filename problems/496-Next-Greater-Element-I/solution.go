type Stack []int

func (s *Stack) Push(num int) {
	*s = append(*s, num)
}

func (s *Stack) Pop() int {
	result := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return result
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums, queries := nums2, nums1
	nextGreater := make(map[int]int)
	stack := make(Stack, 0)
	for i := 0; i < len(nums); i++ {
		nextGreater[nums[i]] = -1
		if stack.Empty() || stack.Top() >= nums[i] {
			stack.Push(nums[i])
		} else {
			for !stack.Empty() && stack.Top() < nums[i] {
				nextGreater[stack.Pop()] = nums[i]
			}
			stack.Push(nums[i])
		}
	}

	result := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		result = append(result, nextGreater[queries[i]])
	}

	return result
}

