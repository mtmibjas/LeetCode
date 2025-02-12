func maximumWealth(accounts [][]int) int {
    w := 0
    for _, a := range accounts {
        for j := 1; j < len(a); j++{
            a[j] = a[j] + a[j-1] 
        }
        if w < a[len(a)-1] {
            w = a[len(a)-1]
        }
    }
    return w
}