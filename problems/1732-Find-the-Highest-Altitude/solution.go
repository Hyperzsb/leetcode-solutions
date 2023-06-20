func largestAltitude(gain []int) int {
    result, sum := 0, 0
    for i := range gain {
        sum += gain[i]

        if sum > result {
            result = sum
        }
    }

    return result
}

