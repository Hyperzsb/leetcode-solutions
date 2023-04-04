func search(target int, arr []int) int {
	if target <= arr[0] {
		return 0
	}

	if target > arr[len(arr)-1] {
		return len(arr)
	}

	start, end := 0, len(arr)-1
	for start < end {
		half := (start + end) / 2

		if target <= arr[half] {
			end = half
		} else {
			start = half + 1
		}
	}

	return (start + end) / 2
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(a, b int) bool {
		return potions[a] < potions[b]
	})

	result := make([]int, len(spells))
	for idx, spell := range spells {
		target := 0
		if success%int64(spell) == 0 {
			target = int(success / int64(spell))
		} else {
			target = int(success/int64(spell)) + 1
		}

		position := search(target, potions)
		result[idx] = len(potions) - position
	}

	return result
}

