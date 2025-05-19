func maximumValue(strs []string) int {
    max := 0
    l := 0
    for _, str := range strs {
        n , err := strconv.Atoi(str)
        if err != nil {
            l = len(str)
        }else {
            l = n
        } 
        if max < l {
            max = l
        }
        
    }   

    return max
}