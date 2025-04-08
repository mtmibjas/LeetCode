func isCovered(ranges [][]int, left int, right int) bool {
    diff := make([]int, 52) 

    for _, r := range ranges {
        diff[r[0]]++
        diff[r[1]+1]--
    }

    sum := 0
    for i := 1; i <= 50; i++ {
        sum += diff[i]
        if i >= left && i <= right && sum == 0 {
            return false
        }
    }

    return true
}