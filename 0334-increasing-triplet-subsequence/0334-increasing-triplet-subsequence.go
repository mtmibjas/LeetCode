func increasingTriplet(nums []int) bool {


	fmin := math.MaxInt64
	smin := math.MaxInt64
	for j := 0; j < len(nums); j++ {
        fmt.Println(fmin, smin)
		if nums[j] <= fmin {
           fmin = nums[j] 
        }else if nums[j] <= smin {
            smin = nums[j]
        }else {
            return true
        }
	}

	
	return false
}