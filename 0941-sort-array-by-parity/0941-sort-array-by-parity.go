func sortArrayByParity(nums []int) []int {
    even := 0
    for i := 0; i < len(nums); i++{
        if nums[i]%2 == 0 {
            nums[even], nums[i] = nums[i],nums[even]  
            even++
        }
    }
    return nums
}