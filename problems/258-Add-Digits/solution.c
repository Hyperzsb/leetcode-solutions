func addDigits(num int) int {
    sum := 0
    for {
        if num == 0 {
            if sum < 10 {
                break
            }
            num = sum
            sum = 0
        }
        sum += num % 10
        num /= 10
    }

    return sum
}

