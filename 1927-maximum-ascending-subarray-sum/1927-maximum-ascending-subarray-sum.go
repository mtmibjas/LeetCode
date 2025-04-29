func maxAscendingSum(nums []int) int {
	right := 0
	left := 0
	max := 0
	for left < len(nums) {
		right = left+1
		sum := nums[left]
	
		for right < len(nums) && nums[right-1] < nums[right] {
			sum += nums[right]
			right++
		}
		if max < sum {
			max = sum
		}
		left++
	}

	return max
}