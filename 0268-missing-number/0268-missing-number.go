func missingNumber(nums []int) int {
    sort.Ints(nums)
    if nums[len(nums)-1] != len(nums) {
            return len(nums)
    }
    for i := 0; i <= len(nums); i++ {
       
       if i != nums[i] {
            return i
       }
    }

    return 0
}