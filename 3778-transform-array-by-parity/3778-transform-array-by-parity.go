func transformArray(nums []int) []int {
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i]%2 == 0 {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
		if nums[j]%2 == 0 {
			nums[j] = 0
		} else {
			nums[j] = 1
		}
	}
	sort.Ints(nums)
	return nums
}