func duplicateZeros(arr []int)  {
    i := 0
    r := 0
    m := make([]int, len(arr))
    for i < len(arr) {
          
        if arr[r] == 0 && i < len(arr)-1 {
            m[i] = 0
            m[i+1] = 0
            i += 2
        }else{
            m[i] = arr[r]
            i++
        }
    
        r++
     
    }
    for i := 0; i < len(arr); i++ {
        arr[i] = m[i]
    }
}