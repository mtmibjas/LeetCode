func rotate(nums []int, k int) {
    n := len(nums)
    if n == 0 || k <= 0 {
        return
    }

    k = k % n 
    
    right := append([]int{}, nums[n-k:]...)
    left := append([]int{}, nums[:n-k]...)

    copy(nums[:k], right)
    copy(nums[k:], left)
}