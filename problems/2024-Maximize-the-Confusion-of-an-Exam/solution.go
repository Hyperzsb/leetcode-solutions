func reverse(s string) string {
	t := make([]byte, len(s))
	for i := range s {
		t[len(s)-1-i] = s[i]
	}

	return string(t)
}

func findMax(answerKey string, k int) int {
	cnts := make([]int, 0)

	prevKey, prevLen := answerKey[0], 1
	for i := 1; i < len(answerKey); i++ {
		if answerKey[i] != prevKey {
			cnts = append(cnts, prevLen)

			prevKey = answerKey[i]
			prevLen = 1
		} else {
			prevLen++
		}
	}
	cnts = append(cnts, prevLen)

	if len(cnts) == 1 {
		return cnts[0]
	}

	preSums := make([][]int, len(cnts))
	preSums[0] = []int{cnts[0], 0}
	for i := 1; i < len(preSums); i++ {
		preSums[i] = make([]int, 2)
		if i%2 == 0 {
			preSums[i][0] = preSums[i-1][0] + cnts[i]
			preSums[i][1] = preSums[i-1][1]
		} else {
			preSums[i][0] = preSums[i-1][0]
			preSums[i][1] = preSums[i-1][1] + cnts[i]
		}
	}

	fmt.Println(cnts, preSums)

	result := 0

	start, end := 0, 0
	for end < len(preSums) {
		sum := preSums[end][0] - preSums[start][0] + cnts[start]

		if sum <= k {
			length := sum + preSums[end][1] - preSums[start][1]

			if start > 0 {
				length += cnts[start-1]
			}

			if end < len(preSums)-1 {
				length += cnts[end+1]
			}

			if result < length {
				result = length
			}
		} else {
			length := k + preSums[end][1] - preSums[start][1]

			if start > 0 {
				length += cnts[start-1]
			}

			if result < length {
				result = length
			}

			for start <= end && sum > k {
				sum -= cnts[start]
				start += 2
			}
		}

		end += 2
	}

	start, end = 1, 1
	for end < len(preSums) {
		sum := preSums[end][1] - preSums[start][1] + cnts[start]

		if sum <= k {
			length := sum + preSums[end][0] - preSums[start][0]

			if start > 0 {
				length += cnts[start-1]
			}

			if end < len(preSums)-1 {
				length += cnts[end+1]
			}

			if result < length {
				result = length
			}
		} else {
			length := k + preSums[end][0] - preSums[start][0]

			if start > 0 {
				length += cnts[start-1]
			}

			if result < length {
				result = length
			}

			for start <= end && sum > k {
				sum -= cnts[start]
				start += 2
			}
		}

		end += 2
	}

	return result
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	forwardResult := findMax(answerKey, k)

	if reverseResult := findMax(reverse(answerKey), k); reverseResult > forwardResult {
		return reverseResult
	} else {
		return forwardResult
	}
}

