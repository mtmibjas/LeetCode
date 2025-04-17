func findIntersectionValues(nums1 []int, nums2 []int) []int {
	arr := []int{}
	c := 0
	for i := 0; i < len(nums1); i++ {
		if slices.Contains(nums2, nums1[i]) {
			c++
		}

	}
	arr = append(arr, c)
	c = 0
	for i := 0; i < len(nums2); i++ {
		if slices.Contains(nums1, nums2[i]) {
			c++
		}

	}
	arr = append(arr, c)

	return arr
}