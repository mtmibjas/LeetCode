func isCovered(ranges [][]int, left int, right int) bool {
    

    for i := left; i <= right; i++ {
        isFind := false
        for _, r := range ranges {
            if r[0] <= i && r[1] >= i {
                isFind = true
                break
            } 
        }
        if !isFind {
            return false
        }
        // fmt.Println(i)
    }

    return true
}