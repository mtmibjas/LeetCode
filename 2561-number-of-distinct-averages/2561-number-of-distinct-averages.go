func distinctAverages(nums []int) int {
    sort.Ints(nums)
    m := make(map[float64]struct{})
    for i,j := 0, len(nums)-1; i < j; i,j = i+1, j-1{
        m[float64(nums[i]+nums[j])/2]=struct{}{}
    }
    return len(m)
}