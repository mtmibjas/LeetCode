func countPartitions(nums []int) int {
    for i := 1; i < len(nums); i++{
        nums[i] += nums[i-1] 
    }
    left, total, res := 0, nums[len(nums)-1], 0
    for left < len(nums)-1 {
        val := (nums[left]-total+nums[left])
        if val%2 == 0 {
            res++
        }
        left++
    }
    return res
}