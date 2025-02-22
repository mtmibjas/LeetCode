func getSneakyNumbers(nums []int) []int {
    m := make(map[int]int)
    var arr []int
    for _, n := range nums{
         m[n]++
    }
    for v, n := range m {
        if  n > 1{
            arr = append(arr, v)
        }
    }
    return arr
}