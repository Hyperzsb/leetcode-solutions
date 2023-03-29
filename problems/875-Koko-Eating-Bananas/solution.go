func isSufficient(piles []int, h, speed int) bool {
	total := 0
	for i := 0; i < len(piles); i++ {
		if piles[i]%speed == 0 {
			total += piles[i] / speed
		} else {
			total += piles[i]/speed + 1
		}
	}

	return total <= h
}

func minEatingSpeed(piles []int, h int) int {
	bananaSum, bananaMax := 0, 0
	for i := 0; i < len(piles); i++ {
		bananaSum += piles[i]

		if bananaMax < piles[i] {
			bananaMax = piles[i]
		}
	}

	minSpeed := bananaSum / h
	if minSpeed == 0 {
		minSpeed = 1
	}
	maxSpeed := bananaMax

	for minSpeed < maxSpeed {
		halfSpeed := (minSpeed + maxSpeed) / 2

		if isSufficient(piles, h, halfSpeed) {
			maxSpeed = halfSpeed
		} else {
			minSpeed = halfSpeed + 1
		}
	}

	return (minSpeed + maxSpeed) / 2
}

