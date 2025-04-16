func permute(nums []int) [][]int {
       res := [][]int{}
    
    var dfs func(int)
    dfs = func(first int) {
        if first == len(nums) {
            // 复制当前排列到结果中
            tmp := make([]int, len(nums))
            copy(tmp, nums)
            res = append(res, tmp)
            return
        }
        // 遍历所有可能的交换位置
        for i := first; i < len(nums); i++ {
            // 交换元素，生成新的排列组合
            nums[first], nums[i] = nums[i], nums[first]
            // 递归处理下一个位置
            dfs(first + 1)
            // 恢复交换，回溯到之前的状态
            nums[first], nums[i] = nums[i], nums[first]
        }
    }
    
    dfs(0)
    return res
}
