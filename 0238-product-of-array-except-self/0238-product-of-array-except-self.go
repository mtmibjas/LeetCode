func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefix, suffix := make([]int,n), make([]int,n)
	prefix[0], suffix[n-1] = nums[0], nums[n-1]
	for i, j := 1, len(nums)-2; i < len(nums); i, j = i+1, j-1 {
		prefix[i], suffix[j] = prefix[i-1]*nums[i], suffix[j+1]*nums[j]
	}

    for i, _ := range nums{
        if i == 0 {
            nums[i] = suffix[i+1]
        }else if i == n-1 {
            nums[i] = prefix[i-1]
        }else{
             nums[i] = prefix[i-1] * suffix[i+1]
        }
    }
	return nums
}