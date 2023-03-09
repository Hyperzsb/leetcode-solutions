func trips(time *[]int, totalTime int64) int64 {
	var result int64 = 0
	for _, t := range *time {
		result += totalTime / int64(t)
	}

	return result
}

func minimumTime(time []int, totalTrips int) int64 {
	sort.Slice(time, func(a, b int) bool {
		return time[a] < time[b]
	})

	var min int64 = 1
	var max int64 = int64(time[len(time)-1] * totalTrips)

	fmt.Println(min, max)

	for min < max {
		var half int64 = (min + max) / 2

		if trips(&time, half) >= int64(totalTrips) {
			max = half
		} else {
			min = half + 1
		}
	}

	return (min + max) / 2
}

