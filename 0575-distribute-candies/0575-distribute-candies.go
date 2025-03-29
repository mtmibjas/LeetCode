func distributeCandies(candyType []int) int {
    l := len(candyType)/2
    m := make(map[int]struct{})

    for _, num := range candyType {
        m[num] = struct{}{}
    }

    if len(m) < l {
        return len(m)
    }
    return l

}