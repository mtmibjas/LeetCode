func divisorSubstrings(num int, k int) int {
	str := strconv.Itoa(num)
    count := 0
	for i := 0; i <= len(str)-k; i++ {
        var d int
		if i == len(str)-k {
			d, _ = strconv.Atoi(string(str[i:]))
		} else {
			d, _ = strconv.Atoi(string(str[i:i+k]))
		}
        if d < 1 {
            continue
        }
        if num%d == 0 {
            count++
        }
	}
    return count
}