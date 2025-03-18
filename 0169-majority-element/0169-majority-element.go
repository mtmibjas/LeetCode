func majorityElement(nums []int) int {
    
    ele, freq := nums[0], 1

    for _, num := range nums[1:] {
        if freq == 0 {
            ele, freq = num, 1
        }else {
            if ele == num {
                freq++
            }else{
                freq--
            }
        }
    }
    return ele
}