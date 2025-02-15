func removeDuplicates(nums []int) int {
    arr := make([]int, 0)
    
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] != nums[i+1]{
            arr = append(arr, nums[i])
        }
    }
    arr = append(arr, nums[len(nums)-1])
    for i := 0; i < len(arr); i++ {
        nums[i] = arr[i]
    }
    return len(arr)
}