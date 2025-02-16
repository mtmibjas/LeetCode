func thirdMax(nums []int) int {
	nums = sort(nums)
	arr := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			arr = append(arr, nums[i])
		}
	}
    arr = append(arr, nums[len(nums)-1])

    if len(arr) < 3 {
		return arr[len(arr)-1]
	}
    
	return arr[len(arr)-3]
}
func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}