func singleNumber(nums []int) []int {
    m := make(map[int]int)
    
    for _, n := range nums {
        m[n]++
    }

    arr := []int{}
    for k, v := range m {
        if v > 1 {
            continue
        }
        arr = append(arr, k)
    }
    return arr
}