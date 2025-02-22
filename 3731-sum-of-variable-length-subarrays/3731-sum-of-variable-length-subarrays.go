func subarraySum(nums []int) int {
    total := 0
    for i := 0; i < len(nums); i++{
        total += func(start, i int) int{
            sum := 0
            for j := start; j <= i; j++{
                sum += nums[j]
            } 
            return sum
        }(max(0,i-nums[i]), i)
    }
    return total
}