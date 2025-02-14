func checkIfExist(arr []int) bool {
    m := make(map[int]int)
    
    for i, num := range arr{
        m[num] = i
    }
    for i, num := range arr{
        if v, ok := m[num*2]; ok && i != v {
            return true
        }
    }
    return false
}