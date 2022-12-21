func countSegments(s string) int {
	segments := strings.Split(s, " ")
	sum := len(segments)
	for _, val := range segments {
		if len(val) == 0 {
			sum--
		}
	}
	return sum
}

