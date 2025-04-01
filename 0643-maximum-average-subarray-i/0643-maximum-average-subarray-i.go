func findMaxAverage(nums []int, k int) float64 {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	m := math.Inf(-1)
	temp := make([]int, 0)
	temp = append(temp, 0)
	temp = append(temp, nums...)
    fmt.Println(temp)
	for i := 1; i < len(temp)-k+1; i++ {
		fmt.Println(temp[i+k-1], temp[i-1])
		avg := float64(temp[i+k-1]-temp[i-1]) / float64(k)
		if avg > m {
			m = avg
		}
	}
	return m
}