func validate(chars []byte) bool {
	if chars[0] == '0' && len(chars) > 1 {
		return false
	}

	sum := 0
	for i := 0; i < len(chars); i++ {
		sum *= 10
		sum += int(chars[i] - '0')
	}

	if sum < 256 {
		return true
	} else {
		return false
	}
}

func restoreS(part int, raw []byte, current *[]string, result *[]string) {
	if part == 1 {
		if len(raw) < 4 {
			return
		}

		for i := 0; i < 3 && i < len(raw)-3; i++ {
			if validate(raw[0 : i+1]) {
				*current = append(*current, fmt.Sprintf("%s", string(raw[0:i+1])))
				restoreS(2, raw[i+1:], current, result)
				*current = (*current)[:0]
			}
		}
	} else if part == 2 {
		if len(raw) < 3 {
			return
		}

		for i := 0; i < 3 && i < len(raw)-2; i++ {
			if validate(raw[0 : i+1]) {
				*current = append(*current, fmt.Sprintf("%s", string(raw[0:i+1])))
				restoreS(3, raw[i+1:], current, result)
				*current = (*current)[:1]
			}
		}
	} else if part == 3 {
		if len(raw) < 2 {
			return
		}

		for i := 0; i < 3 && i < len(raw)-1; i++ {
			if validate(raw[0 : i+1]) {
				*current = append(*current, fmt.Sprintf("%s", string(raw[0:i+1])))
				restoreS(4, raw[i+1:], current, result)
				*current = (*current)[:2]
			}
		}
	} else {
		if len(raw) > 3 || len(raw) == 0 {
			return
		}

		if validate(raw) {
			*result = append(*result,
				fmt.Sprintf("%s.%s.%s.%s",
					(*current)[0], (*current)[1], (*current)[2], raw))
		}
	}
}

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	current := make([]string, 0, 4)
	restoreS(1, []byte(s), &current, &result)

	return result
}

