func findMaxAverage(nums []int, k int) float64 {

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	m := math.Inf(-1)
	nums = append([]int{0}, nums...)
	for i := 1; i < len(nums)-k+1; i++ {
		avg := float64(nums[i+k-1]-nums[i-1]) / float64(k)
		if avg > m {
			m = avg
		}
	}
	return m
}