func countGoodSubstrings(s string) int {
    count := 0
    for i := 0; i <= len(s)-3; i++ {
        m := [26]int{}
        isDis := false
        for _, l := range string(s[i:i+3]) {
            m[l-'a']++
            if m[l-'a'] > 1 {
                isDis = true
                break
            }
        }
        if isDis {
            continue
        }
        count++
    }
    return count
}