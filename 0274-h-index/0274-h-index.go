func hIndex(citations []int) int {
    sort.Slice(citations, func(i, j int)bool{
        return citations[i] > citations[j]
    })
    m := 0
    for i, v := range citations{
       
        if v > 0 && v >= i+1 {
            m = max(m, i+1)
        }
    }
    return m
}