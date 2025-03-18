func firstMissingPositive(nums []int) int {
	n := len(nums)
	pNums := []int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] && nums[i] > 0 {
			pNums = append(pNums, nums[i])
		}
	}
	if nums[n-1] > 0 {
		pNums = append(pNums, nums[n-1])
	}
	fmt.Println(pNums)
	if len(pNums) == 0 {
		return 1
	}
	for i := 1; i <= pNums[len(pNums)-1]; i++ {
		if i != pNums[i-1] {
			return i
		}
	}
	return pNums[len(pNums)-1] + 1
}