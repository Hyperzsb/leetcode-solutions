func isPalindrome(s []byte) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func backtrack(s []byte, current *[]string, result *[][]string) {
	for i := 0; i < len(s); i++ {
		if s[0] == s[i] && isPalindrome(s[:i+1]) {
			*current = append(*current, fmt.Sprintf("%s", string(s[:i+1])))
			if i == len(s)-1 {
				newPartition := make([]string, len(*current))
				copy(newPartition, *current)
				*result = append(*result, newPartition)
			} else {
				backtrack(s[i+1:], current, result)
			}
			*current = (*current)[:len(*current)-1]
		}
	}
}

func partition(s string) [][]string {
	result := make([][]string, 0, 1)
	current := make([]string, 0, 1)

	backtrack([]byte(s), &current, &result)

	return result
}

