func firstMissingPositive(nums []int) int {
	n := len(nums)
	
    for i := 0; i < n; {
        ci := nums[i]-1
        if ci < n && ci >= 0 && nums[i] != nums[ci] {
             nums[i], nums[ci] = nums[ci], nums[i] 
        }else{
            i++
        }
  
    }
    for i := 0; i < n; i++{
        if nums[i]-1 != i {
            return i+1
        }
    }
	return n+1
}