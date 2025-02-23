func sumOfGoodNumbers(nums []int, k int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if i-k >= 0 && i+k <= len(nums)-1 {
			if nums[i-k] < nums[i] && nums[i+k] < nums[i] {
				sum += nums[i]
			}
		} else if i-k < 0 {
			if nums[i+k] < nums[i] {
				sum += nums[i]
			}
		} else if i+k > len(nums)-1 {
			if nums[i-k] < nums[i] {
				sum += nums[i]
			}
		}else{
            sum += nums[i]
        }
        fmt.Println(sum)
	}

	return sum
}