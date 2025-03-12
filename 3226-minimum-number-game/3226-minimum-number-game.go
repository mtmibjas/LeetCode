func numberGame(nums []int) []int {
    sort.Ints(nums)
    
    for i,j := 0,len(nums)-1; i < j; i,j = i+2, j-2 {
        nums[i], nums[i+1] = nums[i+1], nums[i]
        if j == i+1 {
            continue
        }
        nums[j], nums[j-1] = nums[j-1], nums[j]
        
    }

    return nums
    
}