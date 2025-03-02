func getConcatenation(nums []int) []int {
    arr := make([]int, 2*len(nums))
    for i := 0; i < len(nums); i++{
        arr[i] = nums[i]
    }
    j := len(nums)
    for i := 0; i < len(nums); i++{
        arr[j] = nums[i]
        j++
    }
    return arr
}