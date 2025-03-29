func distributeCandies(candyType []int) int {
   
    m := make(map[int]struct{})
    for _, num := range candyType {
        m[num] = struct{}{}
    }
    l := len(candyType)/2
    lm := len(m)
  
    return min(l, lm)

}