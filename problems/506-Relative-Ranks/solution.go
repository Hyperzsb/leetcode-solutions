type Athlete struct {
	Score int
	Idx   int
}

func findRelativeRanks(score []int) []string {
	athletes := make([]Athlete, len(score))
	for i := 0; i < len(score); i++ {
		athletes[i].Score = score[i]
		athletes[i].Idx = i
	}

	sort.Slice(athletes, func(a, b int) bool {
		return athletes[a].Score > athletes[b].Score
	})

	result := make([]string, len(score))

	result[athletes[0].Idx] = "Gold Medal"

	if len(score) == 1 {
		return result
	}

	result[athletes[1].Idx] = "Silver Medal"

	if len(score) == 2 {
		return result
	}

	result[athletes[2].Idx] = "Bronze Medal"

	for i := 3; i < len(score); i++ {
		result[athletes[i].Idx] = fmt.Sprintf("%v", i+1)
	}

	return result
}

