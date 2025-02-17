func containsDuplicate(nums []int) bool {
    m := make(map[int]int)
    for _, num := range nums{
        m[num]++
        if v, _ := m[num]; v > 1{
            return true
        }
    }

    return false
}