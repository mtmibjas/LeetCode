func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]string, numRows)
	currentRow := 0
	direction := 1

	
	for _, char := range s {
		rows[currentRow] += string(char)
		currentRow += direction

		
		if currentRow == 0 || currentRow == numRows-1 {
			direction *= -1
		}
	}

	return strings.Join(rows, "")
}
