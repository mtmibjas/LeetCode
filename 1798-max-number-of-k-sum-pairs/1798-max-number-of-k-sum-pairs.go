func maxOperations(nums []int, k int) int {
    
    sort.Ints(nums)

    left := 0
    right := len(nums)-1

    count := 0
    for left < right{
        if nums[left]+nums[right] == k {
            count++
            left++
            right--
        }else if nums[left]+nums[right] > k{
            right--
        }else{
            left++
        }
    }
    return count
}