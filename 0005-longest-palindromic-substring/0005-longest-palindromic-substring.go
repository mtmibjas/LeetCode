func longestPalindrome(s string) string {
	longStr := ""

	for i := 0; i < len(s); i++ {
		right := len(s) - 1
		str := s[i:]
        find := false
		for right > i {
			if isValid(str) && len(longStr) < len(str) {
				longStr = str
                find = true
                break
			}
			str = s[i:right]
			right--
		}
        if !find && len(longStr) < len(str){
            longStr = string(str[0])
        }
	}
	return longStr
}

func isValid(str string) bool {
	
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}