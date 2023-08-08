func singleNumber(nums []int) int {
    var result int32 = 0

    for i := 0; i <= 31; i++ {
        var sum int32 = 0
        for j := range nums {
            sum += (int32(nums[j]) >> i) & 1
        }

        result |= (sum % 3) << i
    }

    return int(result)
}

