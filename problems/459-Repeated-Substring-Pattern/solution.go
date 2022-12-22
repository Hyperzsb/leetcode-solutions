func check(s string, t string) bool {
	if len(s) == len(t) {
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(t)/len(s); i++ {
		if !check(s, string(t[i*len(s):(i+1)*len(s)])) {
			return false
		}
	}
	return true
}

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if len(s)%i == 0 {
			if check(string(s[:i]), s) {
				return true
			}
		}
	}

	return false
}

