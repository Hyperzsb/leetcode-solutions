type Queue []int

func (q *Queue) Push(n int) {
	*q = append(*q, n)
}

func (q *Queue) Pop() {
	*q = (*q)[1:]
}

func (q *Queue) Front() int {
	return (*q)[0]
}

func minJumps(arr []int) int {
	newArr := make([]int, 1)
	newArr[0] = arr[0]
	preIdx := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[preIdx] {
			continue
		}

		if i-preIdx > 1 {
			newArr = append(newArr, arr[preIdx])
		}

		newArr = append(newArr, arr[i])
		preIdx = i
	}

	if preIdx < len(arr)-1 {
		newArr = append(newArr, arr[preIdx])
	}

	arr = newArr

	visited := make([]bool, len(arr))
	visited[0] = true

	classVisited := make(map[int]bool)

	nodeQ := make(Queue, 0)
	nodeQ.Push(0)

	step := 0
	for len(nodeQ) > 0 {
		length := len(nodeQ)
		for i := 0; i < length; i++ {
			curr := nodeQ.Front()
			nodeQ.Pop()

			if curr == len(arr)-1 {
				return step
			}

			if curr > 0 && !visited[curr-1] {
				nodeQ.Push(curr - 1)
				visited[curr-1] = true
			}

			if curr < len(arr)-1 && !visited[curr+1] {
				nodeQ.Push(curr + 1)
				visited[curr+1] = true
			}

			if !classVisited[arr[curr]] {
				for j := 0; j < len(arr); j++ {
					if arr[j] == arr[curr] && !visited[j] {
						nodeQ.Push(j)
						visited[j] = true
					}
				}

				classVisited[arr[curr]] = true
			}
		}
		step++
	}

	return step
}

