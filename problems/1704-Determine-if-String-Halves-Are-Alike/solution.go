func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)

	count := 0
	for i := 0; i < len(s)/2; i++ {
		if s[i] == 'a' || s[i] == 'A' || s[i] == 'e' || s[i] == 'E' ||
			s[i] == 'i' || s[i] == 'I' || s[i] == 'o' || s[i] == 'O' ||
			s[i] == 'u' || s[i] == 'U' {
			count++
		}
		if s[len(s)-1-i] == 'a' || s[len(s)-1-i] == 'A' ||
			s[len(s)-1-i] == 'e' || s[len(s)-1-i] == 'E' ||
			s[len(s)-1-i] == 'i' || s[len(s)-1-i] == 'I' ||
			s[len(s)-1-i] == 'o' || s[len(s)-1-i] == 'O' ||
			s[len(s)-1-i] == 'u' || s[len(s)-1-i] == 'U' {
			count--
		}
	}

	return count == 0
}

