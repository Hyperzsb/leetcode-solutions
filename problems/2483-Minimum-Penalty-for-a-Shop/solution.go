func bestClosingTime(customers string) int {
    preSums := make([][]int, len(customers)+1)
    for i := range preSums {
        preSums[i] = make([]int, 2)
    }

    for i := 1; i <= len(customers); i++ {
        if customers[i-1] == 'Y' {
            preSums[i][0] = preSums[i-1][0] + 1
            preSums[i][1] = preSums[i-1][1]
        } else {
            preSums[i][0] = preSums[i-1][0]
            preSums[i][1] = preSums[i-1][1] + 1
        }
    }

    fmt.Println(preSums)

    min, idx := preSums[len(customers)][0], 0
    for i := 1; i <= len(customers); i++ {
        if preSums[i][1]+preSums[len(customers)][0]-preSums[i][0] < min {
            min = preSums[i][1]+preSums[len(customers)][0]-preSums[i][0]
            idx = i
        } 
    }

    return idx
}

