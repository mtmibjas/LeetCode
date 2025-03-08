func uniqueOccurrences(arr []int) bool {
    m := make(map[int]int)
    for _, num := range arr {
        m[num]++
    }
    for i, num := range m {
        for j, n := range m {
            if num == n && i != j{
                return false
            }
        }    
    }
    return true
}