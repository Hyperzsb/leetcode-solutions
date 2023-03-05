func compress(chars []byte) int {
	cIdx, rIdx := 0, 0

	for rIdx < len(chars) {
		char, count := chars[rIdx], 0
		for rIdx < len(chars) && chars[rIdx] == char {
			rIdx++
			count++
		}

		chars[cIdx] = char
		cIdx++

		if count == 1 {
			continue
		}

		nIdx := cIdx
		for count > 0 {
			chars[nIdx] = byte(count%10 + '0')
			nIdx++
			count /= 10
		}

		for i := 0; i < (nIdx-cIdx)/2; i++ {
			chars[cIdx+i] ^= chars[nIdx-1-i]
			chars[nIdx-1-i] ^= chars[cIdx+i]
			chars[cIdx+i] ^= chars[nIdx-1-i]
		}

		cIdx = nIdx
	}

	return cIdx
}

