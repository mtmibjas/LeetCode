func minElement(nums []int) int {
	min := 99999

	for i := 0; i < len(nums); i++ {
		s := nums[i] 
		if nums[i] > 9 {
            d := nums[i] 
            s = 0
			for d > 0 {
				r := d % 10
				s += r
				d = d / 10
			}
            
		} 

		if s < min {
			min = s
		}
	}
	return min
}