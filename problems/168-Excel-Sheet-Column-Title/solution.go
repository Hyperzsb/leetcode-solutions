func convertToTitle(columnNumber int) string {
    result := ""
    for columnNumber > 0 {
        remain := columnNumber % 26
        if remain == 0 {
            remain = 26
        }

        result = string('A' + remain - 1) + result 

        columnNumber -= remain
        columnNumber /= 26
    }

    return result
}

