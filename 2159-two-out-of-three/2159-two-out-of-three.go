import "slices"
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    arr := make([]int, 0)
    arr = append(arr, nums1...)
    arr = append(arr, nums3...)
    arr = append(arr, nums3...)
    res := make([]int, 0)
    for _, num := range arr {
        if slices.Contains(res, num) {
            continue
        }
        c := 0
        if slices.Contains(nums1, num) {
            c++
        }
        if slices.Contains(nums2, num) {
            c++
        }
        if slices.Contains(nums3, num) {
            c++
        }
        if c > 1 {
            res = append(res, num)
        }
    }
    return res
}