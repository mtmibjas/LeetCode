func maximumStrongPairXor(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i]-nums[j]))) <= min(nums[i], nums[j]) {
				c := xor(nums[i], nums[j])
				if max < c {
					max = c
				}
			}
		}
	}

	return max
}



func xor(i, j int) int {
	
	return i ^ j
}


