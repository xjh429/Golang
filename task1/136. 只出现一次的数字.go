func singleNumber(nums []int) int {
    res := 0
    for _, num := range nums {
        res ^= num // 异或运算：相同数字异或为 0，最终留下唯一的数
    }
    return res
}
