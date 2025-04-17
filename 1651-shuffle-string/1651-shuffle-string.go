func restoreString(s string, indices []int) string {
	
	
	arr := make([]byte, len(indices))
	for i := 0; i < len(indices); i++ {
		arr[indices[i]] = s[i]
	}
	return string(arr)
}