func reverseDegree(s string) int {

    r := 0
    for i, c := range s {
        d := c-'a'
        r += ((i+1)*(26-int(d)))
    }
    return r
}