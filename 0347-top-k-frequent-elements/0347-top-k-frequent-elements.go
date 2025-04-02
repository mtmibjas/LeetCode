type Num struct{
    Val int
    Key int
}
func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)
    for _, n := range nums{
        m[n]++
    }
    arr := make([]Num, len(m))
    index := 0
    for i, v := range m {
      arr[index] =  Num{Val:v, Key:i}
      index++
    }  
    
    sort.Slice(arr, func(i, j int) bool{
        return arr[i].Val > arr[j].Val
    })

   
    for i := 0; i < k; i++{
        nums[i] = arr[i].Key
    }

    return nums[:k]

}