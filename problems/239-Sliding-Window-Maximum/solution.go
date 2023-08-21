type Deque []int

func (dq *Deque) PushBack(n int) {
    *dq = append(*dq, n)
}

func (dq *Deque) PopBack() {
    *dq = (*dq)[:len(*dq)-1]
}

func (dq *Deque) Back() int {
    return (*dq)[len(*dq)-1]
}

func (dq *Deque) PushFront(n int) {
    newDq := make([]int, 1)
    newDq[0] = n
    newDq = append(newDq, (*dq)...)

    *dq = newDq
}

func (dq *Deque) PopFront() {
    *dq = (*dq)[1:]
}

func (dq *Deque) Front() int {
    return (*dq)[0]
}

func maxSlidingWindow(nums []int, k int) []int {
    result := make([]int, 0)

    if len(nums) == 0 || k == 0 {
        return result
    }

    dq := make(Deque, 0)

    for i := range nums {
        if len(dq) > 0 && i - k + 1 > dq.Front() {
            dq.PopFront()
        }

        for len(dq) > 0 && nums[i] > nums[dq.Back()] {
            dq.PopBack()
        }

        dq.PushBack(i)

        if i >= k-1 {
            result = append(result, nums[dq.Front()])
        }
    }

    return result
}

