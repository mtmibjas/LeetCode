func longestOnes(nums []int, k int) int {
    l, m := 0, 0

    for i := 0; i < len(nums); i++{
        m += (1-nums[i])
        if m > k{
            m -= (1-nums[l])
            l++
        }
    }

    return len(nums)-l
}