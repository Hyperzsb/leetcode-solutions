func getCnt(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return n/2 + 1
	}
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			continue
		}

		if i > 0 {
			flowerbed[i-1] = 1
		}

		if i < len(flowerbed)-1 {
			flowerbed[i+1] = 1
			i++
		}
	}

	cnt, sum := 0, 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			sum += getCnt(cnt)
			cnt = 0
		} else {
			cnt++
		}
	}

	sum += getCnt(cnt)

	return sum >= n
}

