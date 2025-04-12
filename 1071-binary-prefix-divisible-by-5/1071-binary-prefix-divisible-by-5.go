
func prefixesDivBy5(nums []int) []bool {
    res := make([]bool, len(nums))
    var num int

    for i, bit := range nums {
        num = (num*2 + bit) % 5
        res[i] = num == 0
    }

    return res

}

