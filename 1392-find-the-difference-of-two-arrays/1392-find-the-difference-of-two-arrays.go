func findDifference(nums1 []int, nums2 []int) [][]int {
    arr := make([][]int, 2)

    for _, num := range nums1{
        if !slices.Contains(nums2, num) && !slices.Contains(arr[0], num){
            arr[0] = append(arr[0], num)
        }
    }

    for _, num := range nums2{
        if !slices.Contains(nums1, num) && !slices.Contains(arr[1], num){
            arr[1] = append(arr[1], num)
        }
    }

    return arr
}