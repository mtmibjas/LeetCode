func countKConstraintSubstrings(s string, k int) int {
	
	count := 0

	for i := 0; i < len(s); i++ {
		count1, count2 := 0, 0
        if string(s[i]) == "1" {
			count1++
            count++
		}
		if string(s[i]) == "0" {
			count2++
            count++
		}
		left := i + 1
		for left < len(s) {
			if string(s[left]) == "1" {
				count1++
			}
			if string(s[left]) == "0" {
				count2++
			}
            if count1 > k && count2 > k {
                break
            }
            count++
			left++
		}
	}
	return count
}