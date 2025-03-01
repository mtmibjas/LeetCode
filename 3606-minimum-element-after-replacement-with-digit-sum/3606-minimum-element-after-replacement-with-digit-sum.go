func minElement(nums []int) int {
    min := 99999

    for i := 0; i < len(nums); i++{
        d,s := nums[i],0

        for d > 0 {
            r := d%10
            s += r
            d = d/10
        }
        if s < min {
            min = s
        }
    }
    return min
}