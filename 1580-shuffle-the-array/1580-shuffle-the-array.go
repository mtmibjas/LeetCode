func shuffle(nums []int, n int) []int {
    
    i := 0
    arr := make([]int, len(nums))
    for i/2 < n {
        j := i/2
        arr[i],arr[i+1] = nums[j], nums[j+n]
        i += 2
    }
    return arr
}