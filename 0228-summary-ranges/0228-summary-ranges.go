func summaryRanges(nums []int) []string {
	start, end := 0, 0
	arr := make([]string, 0)
	for end < len(nums) {
        if end == len(nums)-1 {
            str := strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end])
			if start == end {
				str = strconv.Itoa(nums[end])
			}
			arr = append(arr, str)
            end++
			continue
        }
		if nums[end]+1 == nums[end+1] {
			end++
			continue
		} else {
			str := strconv.Itoa(nums[start]) + "->" + strconv.Itoa(nums[end])
			if start == end {
				str = strconv.Itoa(nums[end])
			}
			arr = append(arr, str)
			end++
			start = end
		}

	}
	if end == len(nums)-1 {
		str := strconv.Itoa(nums[end])
		arr = append(arr, str)

	}
	return arr
}