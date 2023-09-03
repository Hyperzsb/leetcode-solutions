func format(s, e int, words []string, maxWidth int) string {
    width := 0
    for i := s; i <= e; i++ {
        width += len(words[i])
    }

    line, spaces := "", maxWidth - width
    if s == e {
        line += words[s]
        for i := width; i < maxWidth; i++ {
            line += " "
        }
    } else if spaces%(e-s) == 0 {
        for i := s; i < e; i++ {
            line += words[i]
            for j := 0; j < spaces/(e-s); j++ {
                line += " "
            }
        }

        line += words[e]
    } else {
        for i := s; i < s+spaces%(e-s); i++ {
            line += words[i]
            for j := 0; j < spaces/(e-s)+1; j++ {
                line += " "
            }
        }

        for i := s+spaces%(e-s); i < e; i++ {
            line += words[i]
            for j := 0; j < spaces/(e-s); j++ {
                line += " "
            }
        }

        line += words[e]
    }

    return line
}

func fullJustify(words []string, maxWidth int) []string {
    result := make([]string, 0)

    idx := 0
    for idx < len(words) {
        s, e, width := idx, idx+1, len(words[idx])
        for e < len(words) && width <= maxWidth {
            width += 1 + len(words[e])
            e++
        }

        if width > maxWidth {
            e--
        }

        if e == len(words) {
            line, width := words[s], len(words[s])
            
            for s+1 < e {
                line += " " + words[s+1]
                width += 1 + len(words[s+1])
                s++
            }

            for i := width; i < maxWidth; i++ {
                line += " "
            }

            result = append(result, line)
        } else {
            line := format(s, e-1, words, maxWidth)
            result = append(result, line)
        }

        idx = e
    }

    return result
}

