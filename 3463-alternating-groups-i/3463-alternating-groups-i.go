func numberOfAlternatingGroups(colors []int) int {
    count := 0
    n := colors
    for i := 0; i < len(n); i++ {
        if i > len(n)-2 {
            if n[i] == n[1] && n[i] != n[0] {
                count++
            }
        }else if i > len(n)-3 {
            if n[i] == n[0] && n[i] != n[i+1] {
                count++
            }
        }else{
            if n[i] == n[i+2] && n[i] != n[i+1] {
                count++
            }
        }

    } 
    return count
}