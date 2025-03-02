func getConcatenation(nums []int) []int {
    arr := make([]int, 0)
    arr = append(arr, nums...)
    arr = append(arr,nums...)
    return arr
}