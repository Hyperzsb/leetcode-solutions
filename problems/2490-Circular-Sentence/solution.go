func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")

	for i := 0; i < len(words)-1; i++ {
		if words[i][len(words[i])-1] != words[i+1][0] {
			return false
		}
	}

	if words[len(words)-1][len(words[len(words)-1])-1] != words[0][0] {
		return false
	}

	return true
}

