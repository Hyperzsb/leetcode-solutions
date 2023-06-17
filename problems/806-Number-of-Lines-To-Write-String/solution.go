func numberOfLines(widths []int, s string) []int {
    line, width := 1, 0
    for i := 0; i < len(s); {
        if width+widths[s[i]-'a'] > 100 {
            line++
            width = 0
        } else {
            width += widths[s[i]-'a']
            i++
        }
    }

    return []int{line, width}
}

