func duplicateZeros(arr []int)  {
    i := 0
    for i < len(arr)-1 {
        if arr[i] == 0 {
          
            for j := len(arr)-2; j > i ; j-- {
                arr[j+1] = arr[j] 
            }
            arr[i+1] = 0
            i += 2
        }else{
            i++
        }
    }
}