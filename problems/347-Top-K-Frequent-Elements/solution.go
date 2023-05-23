type Num struct {
	Val  int
	Freq int
}

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	frequency := make([]Num, 0, len(numMap))
	for val, freq := range numMap {
		frequency = append(frequency, Num{val, freq})
	}

	sort.Slice(frequency, func(a, b int) bool {
		return frequency[a].Freq > frequency[b].Freq
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = frequency[i].Val
	}

	return result
}

