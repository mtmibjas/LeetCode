func countBits(n int) []int {
	arr := []int{}

    for i := 0; i <= n; i++{
        arr = append(arr, CountNum(i))
    }
   
	return arr
}

func CountNum(n int)int{
    if n == 0 {
        return 0
    }
    count := 0
    for n > 0 {
		if n%2 == 1{
            count++
        }
		n = n / 2
	}
    return count
}