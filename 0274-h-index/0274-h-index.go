func hIndex(citations []int) int {
    n := len(citations)
    sort.Ints(citations)
    m := 0
    for i, v := range citations{
        if v > 0 && n-i <= v {
            m = max(m, n-i)
        }
    }
    return m
}