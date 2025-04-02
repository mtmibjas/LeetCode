type Num struct{
    Val int
    Key int
}
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, n := range nums{
        m[n]++
    }
    arr := []Num{}
    for i, v := range m {
      arr = append(arr, Num{Val:v, Key:i})
    }  
    
    sort.Slice(arr, func(i, j int) bool{
        return arr[i].Val > arr[j].Val
    })

    res := []int{}
    for i := 0; i < k; i++{
        val := arr[i].Key
        res = append(res, val)
    }

    return res

}