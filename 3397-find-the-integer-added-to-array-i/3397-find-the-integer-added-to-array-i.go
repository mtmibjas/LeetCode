func addedInteger(nums1 []int, nums2 []int) int {
	max1 := 0
	max2 := 0
	for i := 0; i < len(nums1); i++ {
        if max1 < nums1[i]{
            max1 = nums1[i]
        }
        if max2 < nums2[i]{
            max2 = nums2[i]
        }
	}
    return max2-max1
}