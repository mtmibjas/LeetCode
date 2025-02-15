func removeDuplicates(nums []int) int {
    m := make(map[int]int)
    n := 0
    for i:=0; i <len(nums); i++ {
        if v, _ := m[nums[i]]; v < 2 {
            m[nums[i]] = v+1
            nums[n] = nums[i]
            n++ 
        } 
    }   
    return len(nums[:n])
}