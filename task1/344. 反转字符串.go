func reverseString(s []byte) {
    left := 0
    right := len(s) - 1
    for left < right {
        // 交换首尾字符
        s[left], s[right] = s[right], s[left]
        // 向中间收缩指针
        left++
        right--
    }
}
