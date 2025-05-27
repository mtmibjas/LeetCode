func maxFreqSum(s string) int {
	vowels := map[rune]int{
		'a': 0,
		'e': 0,
		'i': 0,
		'o': 0,
		'u': 0,
	}
	consonant := make(map[rune]int)
	mv := 0
	mc := 0
	for _, ru := range s {
		if _, ok := vowels[ru]; ok { 
			vowels[ru]++
			if mv < vowels[ru] {
				mv = vowels[ru]
			}
			continue
		}
		consonant[ru]++
		if v := consonant[ru]; v > mc {
			mc = v
		}
	}
	return mc + mv

}