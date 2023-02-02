type Player struct {
	Age   int
	Score int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func bestTeamScore(scores []int, ages []int) int {
	players := make([]Player, len(scores))
	for i := 0; i < len(players); i++ {
		players[i] = Player{ages[i], scores[i]}
	}

	sort.Slice(players, func(a, b int) bool {
		if players[a].Age < players[b].Age {
			return true
		} else if players[a].Age == players[b].Age {
			return players[a].Score < players[b].Score
		} else {
			return false
		}
	})

	result := 0
	dp := make([]int, len(players))
	for i := 0; i < len(players); i++ {
		dp[i] = players[i].Score
		for j := 0; j < i; j++ {
			if players[j].Score <= players[i].Score {
				dp[i] = max(dp[i], dp[j]+players[i].Score)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

