func reverse(s string) string {
    result := ""
    for i := len(s)-1; i >= 0; i-- {
        result += string(s[i])
    }

    return result
}

func reverseWords(s string) string {
    words := strings.Split(s, " ")

    result := ""
    for i := 0; i < len(words)-1; i++ {
        result += reverse(words[i]) + " "
    }
    result += reverse(words[len(words)-1])

    return result
}

