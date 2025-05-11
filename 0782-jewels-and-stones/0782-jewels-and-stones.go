func numJewelsInStones(jewels string, stones string) int {
    c := 0
    for _, l := range stones {
        if strings.Contains(jewels, string(l)){
            c++
        }
    }

    return c
}