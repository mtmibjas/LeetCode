func sumOfGoodNumbers(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i-k >= 0 && nums[i-k] >= nums[i] {
			continue
		} 
        if i+k < len(nums) && nums[i+k] >= nums[i] {
			continue
		}
            sum += nums[i]
        
	}

	return sum
}