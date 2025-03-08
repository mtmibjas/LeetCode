func uniqueOccurrences(arr []int) bool {
    m := make(map[int]int)
    for _, num := range arr {
        m[num]++
    }
    m2 := make(map[int]struct{})
    for _, num := range m {
           m2[num] = struct{}{}
    }
    return len(m)==len(m2)
}