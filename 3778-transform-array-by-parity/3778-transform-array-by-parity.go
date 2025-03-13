func transformArray(nums []int) []int {
	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		transform(i, nums)
        transform(j, nums)

	}
	sort.Ints(nums)
	return nums
}

func transform(j int, nums []int) {
	if nums[j]%2 == 0 {
		nums[j] = 0
	} else {
		nums[j] = 1
	}
}