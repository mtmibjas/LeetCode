func findMaxAverage(nums []int, k int) float64 {

	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	mx := sum
	for i := 1; i < len(nums)-k+1; i++ {
		sum -= nums[i-1]
		sum += nums[i+k-1]
		if sum > mx {
			mx = sum
		}
	}

	return float64(mx) / float64(k)
}