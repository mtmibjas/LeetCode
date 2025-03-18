func majorityElement(nums []int) int {
    m := make(map[int]int)
    for _, num := range nums{
        m[num]++
    }
    max, c := 0, 0
    for k, v := range m {
        if v > c {
            max = k
            c = v 
        }
    }
    return max
}