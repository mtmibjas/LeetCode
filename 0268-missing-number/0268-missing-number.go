func missingNumber(nums []int) int {
    sort.Ints(nums)
    for i := 0; i <= len(nums); i++ {
       if i == len(nums) {
            return i
       }
       if i != nums[i] {
            return i
       }
    }

    return 0
}