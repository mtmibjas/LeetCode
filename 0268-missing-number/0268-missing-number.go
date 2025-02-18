func missingNumber(nums []int) int {
    m := make(map[int]struct{})

    for i := 0; i < len(nums); i++ {
        m[nums[i]]= struct{}{}
    }

    for i := 0; i <= len(nums); i++ {
        if _,ok := m[i]; !ok{
            return i
        }
    }
    return 0
}