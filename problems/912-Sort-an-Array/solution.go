func partition(nums []int, low, high int) int {
	pivot := nums[low]
	i, j := low-1, high+1

	for {
		i++
		for nums[i] < pivot {
			i++
		}

		j--
		for nums[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}

	return j
}

func quicksort(nums []int, low, high int) {
	if low < high {
		pivot := partition(nums, low, high)
		quicksort(nums, low, pivot)
		quicksort(nums, pivot+1, high)
	}
}

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)

	return nums
}

