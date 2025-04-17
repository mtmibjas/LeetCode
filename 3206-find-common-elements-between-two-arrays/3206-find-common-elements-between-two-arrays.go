func findIntersectionValues(nums1 []int, nums2 []int) []int {
	arr := []int{}
	c1, c2 := 0, 0
	for i, j := 0, len(nums2)-1; i < len(nums1) || j >= 0; i, j = i+1, j-1 {
		if i < len(nums1) && slices.Contains(nums2, nums1[i]) {
			c1++
		}
		if  j >= 0 && slices.Contains(nums1, nums2[j]) {
			c2++
		}
	}
	arr = append(arr, c1, c2)

	return arr
}