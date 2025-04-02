func longestSubarray(nums []int) int {
	left, z := 0, 0
    for right := 0; right < len(nums); right++ {
        z += (1-nums[right])
        if z > 1 {
            z -= (1-nums[left])
            left++
        }
    }
	return len(nums)-left-1
}