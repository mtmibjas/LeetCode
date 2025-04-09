func countBits(n int) []int {
	arr := make([]int, n+1)

    for i := 0; i <= n; i++{
        arr[i] = CountNum(i)
    }
   
	return arr
}

func CountNum(n int)int{
    
    count := 0
    for n > 0 {
		count += n & 1
		n >>= 1
	}
    return count
}