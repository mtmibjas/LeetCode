func longestSubarray(nums []int) int {
	left := 0
	right := 0
	m := 0
	for right < len(nums) {
		if nums[left] == 1 && left < right {
			left++
			continue
		}
		if nums[left] == 0 {
			left++
			right = left
			continue
		}
		right = left
		count := 0
		isDeleted := 0
		for isDeleted < 2 && right < len(nums) {
			if nums[right] == 1 {
				count++
			} else {
				isDeleted++
			}
			right++
		}
        if isDeleted == 0 && left == 0 {
            count--
        }
		if m < count {
			m = count
		}
	}
	return m
}