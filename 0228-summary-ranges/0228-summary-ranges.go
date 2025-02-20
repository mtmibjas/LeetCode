func summaryRanges(nums []int) []string {
	start := 0
	arr := make([]string, 0)
	for end := 0; end < len(nums); end++ {
		if end == len(nums)-1 || nums[end]+1 != nums[end+1] {
			if start == end {
				arr = append(arr, strconv.Itoa(nums[end]))
			} else {
				arr = append(arr, strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end]))
			}
			start = end + 1
		}

	}

	return arr
}