func threeSum(nums []int) [][]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}

	arr := make([][]int, 0)
	for j := 0; j < len(nums)-2; j++ {
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}

		left, right := j+1, len(nums)-1
		for left < right {
			sum := nums[j] + nums[left] + nums[right]
			if sum == 0 {

				arr = append(arr, []int{nums[left], nums[j], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}

		}
	}

	return arr
}