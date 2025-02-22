func getSneakyNumbers(nums []int) []int {
    m := make(map[int]int)
    var arr []int
    for _, n := range nums{
        m[n]++
        if v, _ := m[n]; v > 1{
            arr = append(arr, n)
        }
    }
    return arr
}