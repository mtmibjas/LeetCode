func isSubsequence(s string, t string) bool {
	i := 0
   
    for _, w := range t {
       if i < len(s) && w == rune(s[i]) {
            i++
       }
       if i == len(s) {
          return true
       }    
    }
	 
     return i == len(s) 
}