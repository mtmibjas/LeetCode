func maximumLengthSubstring(s string) int {
    max := 0
    for i := 0; i < len(s); i++ {
        right :=i
        m := make(map[byte]int)
        count := 0
        for right < len(s) {
            m[s[right]]++
            if m[s[right]] > 2 {
                break
            }
            right++
            count++
        }

        if max < count {
            max = count
        }
    }

    return max
}