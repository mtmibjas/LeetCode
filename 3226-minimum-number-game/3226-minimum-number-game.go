func numberGame(nums []int) []int {
    sort.Ints(nums)
    
    for i,j := 0,1; j< len(nums); i,j = i+2, j+2 {
        nums[i], nums[j] = nums[j], nums[i]
    }

    return nums
    
}