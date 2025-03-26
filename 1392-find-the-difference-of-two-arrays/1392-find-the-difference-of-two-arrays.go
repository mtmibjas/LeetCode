func findDifference(nums1 []int, nums2 []int) [][]int {
    arr1 := []int{}

    for _, num := range nums1{
        if !slices.Contains(nums2, num) && !slices.Contains(arr1, num){
            arr1 = append(arr1, num)
        }
    }

    arr2 := []int{}

    for _, num := range nums2{
        if !slices.Contains(nums1, num) && !slices.Contains(arr2, num){
            arr2 = append(arr2, num)
        }
    }

    return [][]int{arr1, arr2}
}