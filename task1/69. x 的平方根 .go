func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    left, right := 1, x
    res := 0
    for left <= right {
        mid := left + (right - left) / 2
        // 用除法代替乘法，避免溢出
        if mid <= x / mid {
            res = mid      // 记录当前可能的解
            left = mid + 1 // 尝试更大的值
        } else {
            right = mid -1 // 调整右边界
        }
    }
    return res
}
