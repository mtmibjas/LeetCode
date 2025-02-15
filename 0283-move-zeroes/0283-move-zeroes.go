func moveZeroes(nums []int)  {

    arr := make([]int, 0)
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            arr = append(arr, nums[i])
        } 
    }  
    for i := 0; i < len(nums); i++ {
        if i > len(arr)-1 {
            nums[i] = 0
        }else{
            nums[i] = arr[i]
        }
    }
}