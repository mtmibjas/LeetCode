func rotate(nums []int, k int)  {
   n := len(nums)
   arr := make([]int, 0)
   k = k%n
   arr = append(arr, nums[n-k:]...)
   arr = append(arr, nums[:n-k]...)
    for i,num := range arr {
        nums[i] = num
    }

}