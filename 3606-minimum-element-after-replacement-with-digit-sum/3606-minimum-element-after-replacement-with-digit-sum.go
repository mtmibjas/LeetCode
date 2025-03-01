func minElement(nums []int) int {
	min := 99999

	for i := 0; i < len(nums); i++ {
		 s := 0
		if nums[i] > 9 {
            d := nums[i] 
			for d > 0 {
				r := d % 10
				s += r
				d = d / 10
			}
            
		} else {
            s = nums[i]
		}

		if s < min {
			min = s
		}
	}
	return min
}