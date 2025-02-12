func countKConstraintSubstrings(s string, k int) int {
    ln := 0
    count := 0
    for ln <= len(s)+1 {
        ln++
        for i :=0; i <= len(s)-ln; i++{
            str := ""
            if i+ln > len(s)-1{
                  str = string(s[i:])
            }else{
                  str = string(s[i:i+ln])
            }

            z := 0
            o := 0
            for _, l := range str {
                if string(l) == "1" {
                    o++
                }
                if string(l) == "0" {
                    z++
                }
            }
            if z <= k || o <= k{
                count++
            }
        }
        
    }
    return count
}