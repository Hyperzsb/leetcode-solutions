type Person struct {
	InDegree  int
	OutDegree int
}

func findJudge(n int, trust [][]int) int {
	persons := make([]Person, n)

	for _, t := range trust {
		persons[t[0]-1].OutDegree++
		persons[t[1]-1].InDegree++
	}

	for i := 0; i < n; i++ {
		if persons[i].InDegree == n-1 && persons[i].OutDegree == 0 {
			return i + 1
		}
	}

	return -1
}

