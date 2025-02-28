func duplicateNumbersXOR(nums []int) int {
    xor := 0
    m := make(map[int]int)

    for i := 0; i < len(nums); i++{
        m[nums[i]]++
        if c, _ := m[nums[i]]; c == 2 {
            xor ^= nums[i]
        }
        
    }
    return xor
}