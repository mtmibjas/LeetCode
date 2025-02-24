func maxAdjacentDistance(nums []int) int {

    n := len(nums)
	max := 0
	for i := 0; i < len(nums); i++ {
		val := nums[i]-nums[(i+1)%n]
        if val < 0 {
            val = val * -1
        }
        if max < val {
            max = val
        }
	}
    
	return max
}