func minimumSumSubarray(nums []int, l int, r int) int {
	left := 0
	sum := 99999999
	for left <= len(nums)-l{
		
		for right := l; right < r+1; right++ {
            total := 0
            if left+right > len(nums) {
                continue
            }
			for i := left; i < left+right; i++ {
				total += nums[i]
			}

			if sum > total && total > 0 {
				sum = total
			}
		}
		left++

	}
	if sum == 99999999 {
		return -1
	}
	return sum
}