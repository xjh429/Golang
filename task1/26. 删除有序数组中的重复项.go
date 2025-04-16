func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    slow := 1 // 慢指针，指向下一个唯一元素应存放的位置
    for fast := 1; fast < len(nums); fast++ {
        // 当发现新元素时（与上一个唯一元素不同）
        if nums[fast] != nums[slow-1] {
            nums[slow] = nums[fast] // 将新元素存放到慢指针位置
            slow++                  // 慢指针后移
        }
    }
    return slow // 返回唯一元素的总数
}
