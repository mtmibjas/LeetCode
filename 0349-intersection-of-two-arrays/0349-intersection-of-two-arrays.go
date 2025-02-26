func intersection(nums1 []int, nums2 []int) []int {
    i := 0
    lngarr, shtarr := nums1, nums2
    if len(lngarr) < len(shtarr){
        lngarr, shtarr = nums2, nums1
    }
    arr := make([]int,0)
    for i < len(shtarr){
        if slices.Contains(lngarr, shtarr[i]) && !slices.Contains(arr, shtarr[i]) {
            arr = append(arr, shtarr[i])
        }
        i++
    }
    return arr
}