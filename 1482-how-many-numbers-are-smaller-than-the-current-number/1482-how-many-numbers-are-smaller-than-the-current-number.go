func smallerNumbersThanCurrent(nums []int) []int {
    
    arr := make([]int, len(nums))

    for i, num := range nums{
        count := 0
        for j, n := range nums {
            if i != j && num > n {
                count++
            } 
        }
        arr[i] = count
    }


    return arr
}