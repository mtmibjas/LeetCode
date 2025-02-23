func shuffle(nums []int, n int) []int {
    
    i := 0
    var arr []int
    for i < n {
        arr = append(arr, nums[i], nums[i+n])
        i++
    }
    return arr
}