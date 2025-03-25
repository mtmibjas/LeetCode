func compress(chars []byte) int {

	count := 1
	index := 0
	for i := 1; i < len(chars); i++ {
		if chars[i-1] != chars[i] {
			chars[index] = chars[i-1]
			if count > 1 {
				arr := []byte(strconv.Itoa(count))
				n := 0
				for n < len(arr) {
					index++
					chars[index] = arr[n]
					n++
				}

			}
			count = 0
			index++
		}

		count++
	}
	//fmt.Println(string(chars))
	chars[index] = chars[len(chars)-1]

	if count > 1 {
		arr := []byte(strconv.Itoa(count))
		n := 0
		for n < len(arr) {
			index++
			chars[index] = arr[n]
			n++
		}

	}

	chars = chars[:index+1]
	return len(chars)
}