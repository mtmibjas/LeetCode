func shuffle(nums []int, n int) []int {
    
    i := 0
    arr := make([]int, len(nums))
    for i < n {
        arr[i*2],arr[i*2+1] = nums[i], nums[i+n]
        i++
    }
    return arr
}