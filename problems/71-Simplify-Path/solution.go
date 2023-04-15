func simplifyPath(path string) string {
	directories := make([]string, 0)

	directory := make([]byte, 0)
	slashCnt, dotCnt, charCnt := 0, 0, 0
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			slashCnt++

			if slashCnt < 2 {
				continue
			}

			if charCnt > 0 || dotCnt > 2 {
				directories = append(directories, string(directory))
			} else if dotCnt == 2 {
				if len(directories) > 0 {
					directories = directories[:len(directories)-1]
				}
			}

			slashCnt = 1
			dotCnt = 0
			charCnt = 0
			directory = directory[:0]
		} else if path[i] == '.' {
			directory = append(directory, path[i])
			dotCnt++
		} else {
			directory = append(directory, path[i])
			charCnt++
		}
	}

	if charCnt > 0 || dotCnt > 2 {
		directories = append(directories, string(directory))
	} else if dotCnt == 2 {
		if len(directories) > 0 {
			directories = directories[:len(directories)-1]
		}
	}

	result := ""
	for _, dir := range directories {
		result = fmt.Sprintf("%s/%s", result, dir)
	}

	if len(directories) == 0 {
		result = "/"
	}

	return result
}

