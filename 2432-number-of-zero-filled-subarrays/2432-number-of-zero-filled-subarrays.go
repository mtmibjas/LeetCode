func zeroFilledSubarray(nums []int) int64 {
    nums = append(nums, 100)
	count := 0
    arr := []int{}
	for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
           count++ 
        }else {
            if count != 0 {
                arr = append(arr, count)
            }
            count = 0
        }
	}
    fmt.Println(arr)
    sum := 0
    for _, num := range arr {
        sum += ((num+num*num)/2)
    } 
    return int64(sum) 
}