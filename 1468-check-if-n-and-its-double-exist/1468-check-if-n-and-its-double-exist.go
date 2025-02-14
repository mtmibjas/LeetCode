func checkIfExist(arr []int) bool {
    m := make(map[int]struct{})

    for _, num := range arr{
        _, ok := m[num*2]  
        if ok  {
            return true
        }
        _, ok = m[num/2]
        if ok &&  num%2 == 0 {
            return true
        }
        m[num] = struct{}{}
    }
    return false
}