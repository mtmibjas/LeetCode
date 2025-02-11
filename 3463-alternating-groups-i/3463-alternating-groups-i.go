func numberOfAlternatingGroups(n []int) int {
    count := 0
    for i := 0; i < len(n); i++ {
         if n[(len(n)+i-1)%len(n)] != n[i] && n[i] != n[(i+1)%len(n)] {
			count++
		}

    } 
    return count
}