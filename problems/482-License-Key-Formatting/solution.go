func split(s string, k int) ([]byte, []byte) {
	leading, remaining := make([]byte, 0), make([]byte, 0)

	idx := 0
	for idx < len(s) {
		if s[idx] != '-' {
			if s[idx] >= 'a' && s[idx] <= 'z' {
				remaining = append(remaining, s[idx]-32)
			} else {
				remaining = append(remaining, s[idx])
			}
		}
		idx++
	}

	if len(remaining) == 0 {
		return leading, remaining
	}

	if len(remaining)%k == 0 {
		leading = append(leading, remaining[:k]...)
		remaining = remaining[k:]
	} else {
		leading = append(leading, remaining[:len(remaining)%k]...)
		remaining = remaining[len(remaining)%k:]
	}

	return leading, remaining
}

func licenseKeyFormatting(s string, k int) string {
	result := ""

	leading, remaining := split(s, k)
	result = fmt.Sprintf("%s", string(leading))

	idx := 0
	for idx < len(remaining) {
		result = fmt.Sprintf("%s-%s", result, string(remaining[idx:idx+k]))
		idx += k
	}

	return result
}

