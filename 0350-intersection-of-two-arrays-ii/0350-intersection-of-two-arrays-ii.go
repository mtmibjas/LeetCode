func intersect(nums1 []int, nums2 []int) []int {
    sharr, lnarr := nums1, nums2
    if len(nums1) > len(nums2) {
        sharr, lnarr = nums2, nums1
    }
    i := 0
    arr := make([]int, 0)
    for i < len(sharr){
        if index := slices.Index(lnarr, sharr[i]); index != -1 {
            arr = append(arr, sharr[i])
            lnarr[index] = -1
        }
        i++
    }
    return arr
}