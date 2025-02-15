func sortArrayByParity(nums []int) []int {
    if len(nums) == 1 {
        return nums
    }
    even := 0
    for i := 0; i < len(nums); i++{
        if nums[i]%2 == 0 {
            nums[even], nums[i] = nums[i],nums[even]  
            even++
        }
    }
    return nums
}