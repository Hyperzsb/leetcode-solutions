func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
    aliceSizeMap, bobSizeMap := make(map[int]bool), make(map[int]bool)
    aliceSum, bobSum := 0, 0

    for _, size := range aliceSizes {
        aliceSizeMap[size] = true
        aliceSum += size
    }

    for _, size := range bobSizes {
        bobSizeMap[size] = true
        bobSum += size
    }

    diff := aliceSum - bobSum

    result := make([]int, 2)
    for size, _ := range aliceSizeMap {
        if bobSizeMap[-(diff-2*size)/2] {
            result[0] = size
            result[1] = -(diff-2*size)/2

            break
        }
    }
    
    return result
}

