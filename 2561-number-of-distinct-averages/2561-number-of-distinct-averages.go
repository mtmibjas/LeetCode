func distinctAverages(nums []int) int {
    sort.Ints(nums)

    b := 0
    e := len(nums)-1
    m := make(map[float64]struct{})
    for b < e{
        f := float64(nums[b]+nums[e])/2
        m[f]=struct{}{}
        b++
        e--
    }
    return len(m)
}