func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix := make([]int, n), make([]int, n)
	prefix[0], suffix[n-1] = 1, 1
	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		prefix[i], suffix[j] = prefix[i-1]*nums[i-1], suffix[j+1]*nums[j+1]
	}
	for i, _ := range nums {
		nums[i] = prefix[i] * suffix[i]
	}
	return nums
}