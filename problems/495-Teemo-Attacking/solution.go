func findPoisonedDuration(timeSeries []int, duration int) int {
	result := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		if timeSeries[i]+duration <= timeSeries[i+1] {
			result += duration
		} else {
			result += timeSeries[i+1] - timeSeries[i]
		}
	}
	result += duration

	return result
}

