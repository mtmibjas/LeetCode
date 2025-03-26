func findDifference(nums1 []int, nums2 []int) [][]int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	arr := make([][]int, 2)
	left := 0
	right := 0
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] == nums2[right] {
            left++
            for i := left; i < len(nums1); i++{
                if nums1[left-1] == nums1[left]{
                     left++
                     continue
                }
                break
            }
			
			right++
            for i := right; i < len(nums2); i++{
                if nums2[right-1] == nums2[right]{
                     right++
                     continue
                }
                break
            }
			continue
		} else if nums1[left] < nums2[right] {
			if !slices.Contains(arr[0], nums1[left]) {
				arr[0] = append(arr[0], nums1[left])
			}
			left++
		} else {
			if !slices.Contains(arr[1], nums2[right]) {
				arr[1] = append(arr[1], nums2[right])
			}
			right++
		}
          fmt.Println(left, right)
	}
  
	if left < len(nums1) {
		for _, num := range nums1[left:] {
			if !slices.Contains(arr[0], num) {
				arr[0] = append(arr[0], num)
			}
		}

	}
	if right < len(nums2) {
		for _, num := range nums2[right:] {
			if !slices.Contains(arr[1], num) {
				arr[1] = append(arr[1], num)
			}
		}

	}
	return arr
}