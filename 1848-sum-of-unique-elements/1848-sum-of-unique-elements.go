func sumOfUnique(nums []int) int {

	sort.Ints(nums)
	sum := nums[0]
	is := false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			if !is {
				sum -= nums[i]
			}
			is = true
			continue
		}
		is = false
		sum += nums[i]
	}
	return sum
}