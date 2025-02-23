func shuffle(nums []int, n int) []int {
    
    i := 0
    arr := make([]int, len(nums))
    for i/2 < n {
        arr[i],arr[i+1] = nums[i/2], nums[i/2+n]
        i += 2
    }
    return arr
}