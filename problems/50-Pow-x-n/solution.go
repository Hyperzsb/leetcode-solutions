func myPow(x float64, n int) float64 {
    powers := make(map[int]float64)
    powers[0] = 1
    powers[1] = x

    for i := 2; i <= 32; i++ {
        powers[i] = powers[i-1] * powers[i-1]
    }

    flag := false
    if n < 0 {
        flag = true
        n = -n
    }

    result := powers[0]

    index := 1
    for n > 0 {
        if n&1  == 1 {
            result *= powers[index]
        }

        n >>= 1

        index++
    }

    if flag {
        return 1 / result
    } else {
        return result
    }
}

