func shuffle(nums []int, n int) []int {
    
    i, j := 0, 0
    arr := make([]int, len(nums))
    for j < n {
        arr[i],arr[i+1] = nums[j], nums[j+n]
        i += 2
        j++
    }
    return arr
}